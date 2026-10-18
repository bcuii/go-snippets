//go:build ignore

// for + select: keep handling events until a stop signal arrives.
// Here: do periodic work on every tick, stop after a deadline.
//
// Run: go run ./select/basic/05_for_select.go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	stop := time.After(350 * time.Millisecond)

	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-stop:
			fmt.Println("stop")
			return // break here would only exit the select, not the for
		}
	}
}
