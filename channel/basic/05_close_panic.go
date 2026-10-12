//go:build ignore

// What panics around close:
// send on a closed channel, close a channel twice, close a nil channel.
// Receiving from a closed channel never panics (see 04_receive_ok.go).
//
// To stay safe: only the sender closes, and closes exactly once. A receiver never closes.
//
// Run: go run ./channel/basic/05_close_panic.go
package main

func main() {
	ch := make(chan int, 1)
	close(ch)

	// The program stops at the first panic, so only the first one actually runs.
	ch <- 1   // panic: send on closed channel
	close(ch) // panic: close of closed channel

	var nilCh chan int
	close(nilCh) // panic: close of nil channel
}
