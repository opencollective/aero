package aero

import "sync"

// Parallel starts all functions asynchronously as goroutines and waits until they are completed.
func Parallel(funcs ...func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(funcs))

	for _, fun := range funcs {
		task := fun
		go func() {
			task()
			wg.Done()
		}()
	}

	wg.Wait()
}
