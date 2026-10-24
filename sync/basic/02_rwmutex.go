//go:build ignore

// RWMutex: many readers can hold RLock at the same time; a writer's Lock is exclusive.
// Use it for read-heavy data such as caches.
// Here 3 readers each read for 100ms and all finish at 100ms, not 100/200/300ms.
// The writer has to wait until every reader has unlocked.
//
// Run: go run ./sync/basic/02_rwmutex.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var mu sync.RWMutex
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			mu.RLock()
			defer mu.RUnlock()
			time.Sleep(100 * time.Millisecond) // reading
			fmt.Println(time.Since(start).Round(10*time.Millisecond), "reader", i, "done")
		})
	}

	time.Sleep(10 * time.Millisecond) // let the readers take the lock first
	mu.Lock()                         // writer: waits for all readers
	fmt.Println(time.Since(start).Round(10*time.Millisecond), "writer got the lock")
	mu.Unlock()

	wg.Wait()
}
