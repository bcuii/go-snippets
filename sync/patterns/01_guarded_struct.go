//go:build ignore

// Guarded struct: put the mutex right above the fields it protects, and lock inside every method.
// Callers just call methods and never touch the lock.
// Methods use pointer receivers: a value receiver would copy the struct, lock included.
//
// Run:   go run ./sync/patterns/01_guarded_struct.go
// Check: go run -race ./sync/patterns/01_guarded_struct.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	c := NewCounter()
	var wg sync.WaitGroup
	for range 100 {
		wg.Go(func() { c.Inc("home") })
	}
	wg.Wait()
	fmt.Println("home:", c.Get("home"))
}

type Counter struct {
	mu     sync.Mutex // guards counts
	counts map[string]int
}

func NewCounter() *Counter {
	return &Counter{counts: map[string]int{}}
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key]++ // without the lock: fatal error: concurrent map writes
}

func (c *Counter) Get(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counts[key]
}
