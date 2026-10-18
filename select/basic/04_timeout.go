//go:build ignore

// Timeout: time.After returns a channel that receives once after the duration.
// Put it in a select to stop waiting for a slow channel.
//
// Run: go run ./select/basic/04_timeout.go
package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string) // nobody sends: stands in for a slow call

	select {
	case v := <-result:
		fmt.Println("result:", v)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout")
	}
}
