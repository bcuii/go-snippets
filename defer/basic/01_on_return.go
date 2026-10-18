//go:build ignore

// defer schedules a call to run when the surrounding function returns, on every return path.
// That's why cleanup is written right after acquiring a resource.
//
// Run: go run ./defer/basic/01_on_return.go
package main

import "fmt"

func work(fail bool) {
	fmt.Println("open")
	defer fmt.Println("close") // runs on both return paths below

	if fail {
		fmt.Println("early return")
		return
	}
	fmt.Println("done")
}

func main() {
	work(false)
	fmt.Println("---")
	work(true)
}
