package main

import (
	"fmt"
	"sync"
)

// Container holds a map of counters; we add a Mutex
// to update it concurrently from multiple goroutines
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Mutexes must not be copied, so when passed we do it as a pointer
func (c *Container) inc(name string) {
	// Lock the mutex before accesing counters
	// Unlock it at the end of the function using
	// defer statement
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++

}

func main() {
	c := Container{
		// Note that zero value if a muter ix usable as-is
		// so no initizalization is required
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	// Functions increments a named counter in a loop
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)

	wg.Wait()
	fmt.Println(c.counters)

}
