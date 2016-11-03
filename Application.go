package aero

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"crypto/sha256"

	"encoding/base64"

	"github.com/buaazp/fasthttprouter"
	"github.com/fatih/color"
	cache "github.com/patrickmn/go-cache"
	"github.com/valyala/fasthttp"
)

const (
	gzipCacheDuration = 5 * time.Minute
	gzipCacheCleanup  = 1 * time.Minute
)

// Application represents a single web service.
type Application struct {
	Config   Configuration
	Layout   func(*Context, string) string
	Security struct {
		Key         []byte
		Certificate []byte
	}

	css            string
	cssHash        string
	cssReplacement string
	root           string
	router         *fasthttprouter.Router
	gzipCache      *cache.Cache
	start          time.Time
	requestCount   uint64
}

// New creates a new application.
func New() *Application {
	app := new(Application)
	app.root = ""
	app.router = fasthttprouter.New()
	app.gzipCache = cache.New(gzipCacheDuration, gzipCacheCleanup)
	app.start = time.Now()
	app.showStatistics("/__/")
	app.Layout = func(ctx *Context, content string) string {
		return content
	}
	app.Config.Reset()

	return app
}

// Get registers your function to be called when a certain path has been requested.
func (app *Application) Get(path string, handle Handle) {
	app.router.GET(path, func(fasthttpContext *fasthttp.RequestCtx, params fasthttprouter.Params) {
		ctx := Context{
			App:        app,
			Params:     params,
			requestCtx: fasthttpContext,
			start:      time.Now(),
		}

		response := handle(&ctx)
		ctx.Respond(response)

		atomic.AddUint64(&app.requestCount, 1)
	})
}

// Register calls app.Get for both /route and /_/route
func (app *Application) Register(path string, handle Handle) {
	app.Get("/_"+path, handle)
	app.Get(path, func(ctx *Context) string {
		page := handle(ctx)
		html := app.Layout(ctx, page)
		return strings.Replace(html, "</head><body", app.cssReplacement, 1)
	})
}

// SetStyle ...
func (app *Application) SetStyle(css string) {
	app.css = css

	hash := sha256.Sum256([]byte(css))
	app.cssHash = base64.StdEncoding.EncodeToString(hash[:])
	app.cssReplacement = "<style>" + app.css + "</style></head><body"
}

// Run calls app.Load() and app.Listen().
func (app *Application) Run() {
	app.Load()
	app.Listen()
}

// Load loads the application data from the file system.
func (app *Application) Load() {
	// TODO: ...
}

// Listen starts the server.
func (app *Application) Listen() {
	fmt.Println("Server running on:", color.GreenString("http://localhost:"+strconv.Itoa(app.Config.Ports.HTTP)))

	listener := app.listen()
	app.serve(listener)
}

// listen listens on the specified host and port.
func (app *Application) listen() net.Listener {
	address := ":" + strconv.Itoa(app.Config.Ports.HTTP)

	listener, bindError := net.Listen("tcp", address)

	if bindError != nil {
		panic(bindError)
	}

	return listener
}

// serve serves requests from the given listener.
func (app *Application) serve(listener net.Listener) {
	server := &fasthttp.Server{
		Handler: app.router.Handler,
	}

	if app.Security.Key != nil && app.Security.Certificate != nil {
		serveError := server.ServeTLSEmbed(listener, app.Security.Certificate, app.Security.Key)

		if serveError != nil {
			panic(serveError)
		}
	}

	serveError := server.Serve(listener)

	if serveError != nil {
		panic(serveError)
	}
}