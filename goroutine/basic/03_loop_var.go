//go:build ignore

// Since Go 1.22, each loop iteration has its own variable, so each closure captures a different i.
// The `go` line in go.mod decides this, not the compiler version.
//
// Before Go 1.22, all iterations shared one i. By the time the goroutines ran, the loop had
// usually finished, so they printed 3 3 3. Old code fixed it with `i := i`:
//
//	for i := 0; i < 3; i++ {
//		i := i // new variable per iteration; not needed since Go 1.22
//		go func() { fmt.Println(i) }()
//	}
//
// Run: go run ./goroutine/basic/03_loop_var.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Go(func() {
			fmt.Println(i) // 0, 1, 2 in any order
		})
	}
	wg.Wait()
}
