//go:build ignore

// Data race: goroutines access the same variable without synchronization, and at least one writes.
// counter++ is read, add, write. Two goroutines can read the same old value, so updates get lost.
// Fix: Mutex or atomic (see the sync topic).
//
// Run:    go run ./goroutine/basic/04_data_race.go
// Detect: go run -race ./goroutine/basic/04_data_race.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	for range 100 {
		wg.Go(func() {
			for range 1000 {
				counter++ // unsynchronized read-modify-write
			}
		})
	}
	wg.Wait()
	fmt.Println("counter:", counter, "(expected 100000)")
}
