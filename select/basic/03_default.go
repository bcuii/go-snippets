//go:build ignore

// default runs when no case is ready, so select never blocks.
// Common use: try to receive, or try to send and drop the value if the buffer is full.
//
// Run: go run ./select/basic/03_default.go
package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// Try receive
	select {
	case v := <-ch:
		fmt.Println("got", v)
	default:
		fmt.Println("empty, moving on") // ch has no data
	}

	// Try send
	for i := 1; i <= 2; i++ {
		select {
		case ch <- i:
			fmt.Println("sent", i)
		default:
			fmt.Println("full, dropped", i) // buffer of 1 is already taken
		}
	}
}
