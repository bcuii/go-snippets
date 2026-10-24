//go:build ignore

// sync.Map: a map that is safe for concurrent use without extra locking.
// The docs say it only beats map + Mutex in two cases: keys written once and then only read
// (caches that only grow), or goroutines working on disjoint sets of keys.
// Otherwise use map + Mutex. Keys and values are `any`, so reads need a type assertion.
//
// Run: go run ./sync/basic/07_sync_map.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			m.Store(i, i*i) // concurrent writes are fine
		})
	}
	wg.Wait()

	v, ok := m.Load(2)
	fmt.Println("Load(2):", v, ok)

	actual, loaded := m.LoadOrStore(2, 999) // key exists: keeps 4, returns it
	fmt.Println("LoadOrStore(2, 999):", actual, loaded)

	m.Range(func(k, v any) bool {
		fmt.Println("range:", k, v) // random order
		return true                 // false would stop the iteration
	})
}
