//go:build ignore

// Safe goroutine: start goroutines through a wrapper that gives each one its own recover,
// so one panicking task is logged and the rest of the program keeps running.
// This is the fix for ../basic/04_goroutine.go.
//
// Run: go run ./panic/patterns/02_safe_go.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		safeGo(&wg, func() {
			if i == 2 {
				panic("task 2 failed")
			}
			fmt.Println("task", i, "done")
		})
	}
	wg.Wait()
	fmt.Println("main: still running")
}

func safeGo(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done() // deferred first, so it runs last: after the recover below
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recovered:", r) // real code: log r and debug.Stack()
			}
		}()
		f()
	}()
}
