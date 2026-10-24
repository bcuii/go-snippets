//go:build ignore

// CAS (CompareAndSwap): store the new value only if the current value is still the old one.
// Retry loop: load, compute, CAS. If another goroutine changed it in between, the CAS fails: try again.
// Here 100 goroutines report numbers and highest keeps the largest.
//
// Atomic means indivisible: nobody can see the operation half done or slip in between its steps.
// The loop as a whole is not atomic; only Load and CompareAndSwap are.
// A Mutex keeps everyone else out for the whole read-compare-write.
// CAS lets others slip in, and checks whether they did at the moment of writing.
//
// Run: go run ./sync/basic/06_cas.go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var highest atomic.Int64
	for i := 1; i <= 100; i++ {
		wg.Go(func() {
			updateMax(&highest, int64(i))
		})
	}
	wg.Wait()
	fmt.Println("highest:", highest.Load())
}

func updateMax(highest *atomic.Int64, v int64) {
	for {
		old := highest.Load() // atomic read
		// other goroutines may change highest from here...
		if v <= old {
			return // safe without a lock: highest only grows, so v <= old means v <= any later value
		}
		// ...to here
		if highest.CompareAndSwap(old, v) { // atomic: "still old? then write v" as one step
			return // it was still old at that instant: same as doing the whole update at once
		}
		// CAS failed: someone changed it and nothing was written; retry with the new value
	}
}
