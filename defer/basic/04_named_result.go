//go:build ignore

// A deferred closure can modify named return values:
// return sets err first, then deferred calls run, then the caller gets err.
// Common use: wrap every error path in one place.
//
// Run: go run ./defer/basic/04_named_result.go
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(load("config.json"))
}

func load(name string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("load %s: %w", name, err) // rewrites the value being returned
		}
	}()

	return errors.New("not found") // sets err, then the defer above runs
}
