//go:build ignore

// WaitGroup waits for a group of goroutines. It is a counter:
// Add increments it, Done decrements it, Wait blocks until it reaches 0.
// Since Go 1.25, wg.Go(f) does the Add and Done for you.
//
// Run: go run ./goroutine/basic/02_waitgroup.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	classic()
	withGo()
}

func classic() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // before go: inside the goroutine, Wait might run first and see 0
		go func() {
			defer wg.Done()
			fmt.Println("classic: task", i)
		}()
	}
	wg.Wait()
}

func withGo() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Go(func() { // Go 1.25+: Add(1) before starting, Done when the func returns
			fmt.Println("wg.Go: task", i)
		})
	}
	wg.Wait()
}
