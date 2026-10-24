//go:build ignore

// go f() starts f in a new goroutine and returns immediately: the caller doesn't wait.
// When main returns, the program exits and every other goroutine is killed.
//
// Run: go run ./goroutine/basic/01_go_keyword.go
package main

import (
	"fmt"
	"time"
)

func main() {
	go work()
	fmt.Println("main: started work, not waiting")
	// main returns here, so work is killed before it finishes
}

func work() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("work: done") // never printed
}
