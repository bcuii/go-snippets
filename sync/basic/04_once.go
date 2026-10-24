//go:build ignore

// Once: run initialization exactly once, even when many goroutines ask for it at the same time.
// The others wait until it finishes. Since Go 1.21, sync.OnceValue also returns the result.
//
// Run: go run ./sync/basic/04_once.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	classic()
	withOnceValue()
}

func classic() {
	var once sync.Once
	var config string
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			once.Do(func() {
				fmt.Println("classic: loading config") // printed once
				config = "loaded"
			})
			fmt.Println("classic: goroutine", i, "sees", config)
		})
	}
	wg.Wait()
}

func withOnceValue() {
	loadConfig := sync.OnceValue(func() string { // Go 1.21+
		fmt.Println("OnceValue: loading config") // printed once
		return "loaded"
	})
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			fmt.Println("OnceValue: goroutine", i, "sees", loadConfig())
		})
	}
	wg.Wait()
}
