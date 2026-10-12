//go:build ignore

// Buffered channel: sends don't block until the buffer is full (a 3rd send would block).
//
// Run: go run ./channel/basic/02_buffered.go
package main

import "fmt"

func main() {
	buf := make(chan int, 2)
	buf <- 11
	buf <- 33
	fmt.Println("buffered: len =", len(buf), "cap =", cap(buf))
}
