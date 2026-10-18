//go:build ignore

// select waits on several channel operations at once and runs the one that is ready.
// If none is ready, it blocks.
//
// Run: go run ./select/basic/01_multiplex.go
package main

import "fmt"

func main() {
	a := make(chan string, 1)
	b := make(chan string, 1)
	b <- "hello"

	select {
	case v := <-a:
		fmt.Println("from a:", v)
	case v := <-b:
		fmt.Println("from b:", v) // only b has data, so this case runs
	}

	// Both are empty now: select blocks forever.
	select {
	case <-a:
	case <-b:
	}
	// never reached
}
