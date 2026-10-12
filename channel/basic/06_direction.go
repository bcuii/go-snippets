//go:build ignore

// Directional channel types: chan<- T can only send, <-chan T can only receive.
// The compiler enforces it, so for example a receiver can't close the channel.
//
// Run: go run ./channel/basic/06_direction.go
package main

import "fmt"

func producer(out chan<- int) {
	out <- 11
	out <- 33
	close(out)
	// <-out // compile error: cannot receive from send-only channel
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Println("got", v)
	}
	// in <- 1   // compile error: cannot send to receive-only channel
	// close(in) // compile error: cannot close receive-only channel
}

func main() {
	ch := make(chan int, 2) // bidirectional
	producer(ch)            // implicitly converted to chan<- int
	consumer(ch)            // implicitly converted to <-chan int
}
