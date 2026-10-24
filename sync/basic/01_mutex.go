//go:build ignore

// Mutex: only one goroutine at a time runs the code between Lock and Unlock (the critical section).
// This fixes the lost updates in goroutine/basic/04_data_race.go.
//
// Run:   go run ./sync/basic/01_mutex.go
// Check: go run -race ./sync/basic/01_mutex.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0
	for range 100 {
		wg.Go(func() {
			for range 1000 {
				mu.Lock()
				counter++   // critical section
				mu.Unlock() // in a loop, unlock directly: a defer would wait for the function to return
			}
		})
	}
	wg.Wait()
	fmt.Println("counter:", counter, "(expected 100000)")
}
