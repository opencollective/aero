package aero

// This list includes all the common header keys
// and values used in the http server code.
const (
	cacheControlHeader            = "Cache-Control"
	cacheControlAlwaysValidate    = "must-revalidate"
	cacheControlMedia             = "public, max-age=13824000"
	contentTypeOptionsHeader      = "X-Content-Type-Options"
	contentTypeOptions            = "nosniff"
	xssProtectionHeader           = "X-XSS-Protection"
	xssProtection                 = "1; mode=block"
	etagHeader                    = "ETag"
	contentTypeHeader             = "Content-Type"
	contentTypeHTML               = "text/html; charset=utf-8"
	contentTypeCSS                = "text/css; charset=utf-8"
	contentTypeJavaScript         = "application/javascript; charset=utf-8"
	contentTypeJSON               = "application/json; charset=utf-8"
	contentTypePlainText          = "text/plain; charset=utf-8"
	contentTypeEventStream        = "text/event-stream; charset=utf-8"
	contentTypeSVG                = "image/svg+xml"
	contentEncodingHeader         = "Content-Encoding"
	contentEncodingGzip           = "gzip"
	acceptEncodingHeader          = "Accept-Encoding"
	contentLengthHeader           = "Content-Length"
	ifNoneMatchHeader             = "If-None-Match"
	referrerPolicyHeader          = "Referrer-Policy"
	referrerPolicySameOrigin      = "no-referrer"
	strictTransportSecurityHeader = "Strict-Transport-Security"
	strictTransportSecurity       = "max-age=31536000; includeSubDomains; preload"
	contentSecurityPolicyHeader   = "Content-Security-Policy"
	forwardedForHeader            = "X-Forwarded-For"
	realIPHeader                  = "X-Real-Ip"
)
