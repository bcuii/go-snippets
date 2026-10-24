//go:build ignore

// atomic: an indivisible read, write or add on one variable, done by a CPU instruction, no lock.
// Use the typed API (atomic.Int64, atomic.Bool, atomic.Pointer[T], Go 1.19+).
// It only protects a single variable; for several fields that must change together, use a Mutex.
//
// Run:   go run ./sync/basic/05_atomic.go
// Check: go run -race ./sync/basic/05_atomic.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter atomic.Int64
	for range 100 {
		wg.Go(func() {
			for range 1000 {
				counter.Add(1) // read, add, write as one step
			}
		})
	}
	wg.Wait()
	fmt.Println("counter:", counter.Load(), "(expected 100000)")
}
