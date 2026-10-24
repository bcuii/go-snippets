//go:build ignore

// Must: panic on an error that can only come from a bug, e.g. a typo in a constant.
// Typically used for package-level vars, so the program fails at startup instead of later.
// Input that can legitimately be wrong (from users, files, network) uses the error-returning version.
//
// Run: go run ./panic/patterns/01_must.go
package main

import (
	"fmt"
	"regexp"
)

var digits = regexp.MustCompile(`^\d+$`) // constant pattern: an error here is a bug

func main() {
	fmt.Println(digits.MatchString("123"))

	// User input: handle the error.
	if _, err := regexp.Compile(`[`); err != nil {
		fmt.Println("bad user pattern:", err)
	}

	regexp.MustCompile(`[`) // panic: same bad pattern through Must
}
