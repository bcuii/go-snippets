//go:build ignore

// Goroutines are cheap: each starts with a small stack (a few KB) that grows as needed,
// while an OS thread usually reserves MBs. Running 100,000 of them is normal.
//
// The scheduler (GMP) runs many goroutines (G) on a few OS threads (M).
// A thread needs a logical processor (P) to run Go code, and there are GOMAXPROCS Ps.
//
// Run: go run ./goroutine/basic/05_lightweight.go
package main

import (
	"fmt"
	"runtime"
	"runtime/pprof"
)

func main() {
	const n = 100_000
	block := make(chan struct{})
	for range n {
		go func() { <-block }() // keep every goroutine alive at the same time
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("goroutines:   ", runtime.NumGoroutine())
	fmt.Println("stack per G:  ", m.StackInuse/n, "bytes")
	fmt.Println("OS threads:   ", pprof.Lookup("threadcreate").Count()) // M: only a handful
	fmt.Println("GOMAXPROCS:   ", runtime.GOMAXPROCS(0))                // P: defaults to CPU count
}
