//go:build ignore

// Unbuffered channel: a send blocks until someone receives.
//
// Run: go run ./channel/basic/01_unbuffered.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		time.Sleep(3000 * time.Millisecond)
		<-ch // the receiver shows up later
	}()
	start := time.Now()
	ch <- "hello"
	fmt.Println("unbuffered: send blocked for", time.Since(start).Round(100*time.Millisecond))
}
