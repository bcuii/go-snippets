//go:build ignore

// Idle timeout: stop when no message arrives for 200ms, e.g. closing an idle connection.
// time.After inside the loop is intended here: each select starts a new timer, so every message resets the countdown.
// For one deadline over the whole loop, create the timer (or a ctx) once, before the loop.
//
// Run: go run ./select/patterns/03_idle_timeout.go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	msgs := make(chan string)
	go func() {
		for range 3 { // a message every 100ms, then silence
			time.Sleep(100 * time.Millisecond)
			msgs <- "msg"
		}
	}()

	for {
		select {
		case m := <-msgs:
			fmt.Println(time.Since(start).Round(10*time.Millisecond), "got", m)
		case <-time.After(200 * time.Millisecond):
			fmt.Println(time.Since(start).Round(10*time.Millisecond), "idle for 200ms: close")
			return
		}
	}
}
