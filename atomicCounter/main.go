package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// Represents an always positive number
	var ops atomic.Uint64

	// WaitGroup to wait for all the goroutines to execute
	var wg sync.WaitGroup

	// start 50 goroutines taht each will increment the counter 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	// Wait until the goroutines are done
	wg.Wait()

	// no go routines are writing ops but using the Load()
	// it's safe to atomically read a value even while other goroutines
	// are (atomically) updating it
	fmt.Println("ops:", ops.Load())

}
