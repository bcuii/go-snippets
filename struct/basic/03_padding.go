//go:build ignore

// Memory alignment: each field starts at an offset that is a multiple of its alignment (8 for int64),
// so the compiler adds unused bytes (padding) between fields.
// Field order changes the struct size. It matters for large slices of structs.
//
// Size and alignment on 64-bit (unsafe.Sizeof, unsafe.Alignof):
//
//	type                    size  align
//	bool, int8, byte          1     1
//	int16                     2     2
//	int32, float32, rune      4     4
//	int, int64, float64, *T   8     8
//	map, chan, func           8     8   (a single pointer)
//	string                   16     8   (pointer + length)
//	any, interfaces          16     8   (type + data pointer)
//	[]T                      24     8   (pointer + length + capacity)
//	[3]int16                  6     2   (an array aligns like its element)
//
// Run: go run ./struct/basic/03_padding.go
package main

import (
	"fmt"
	"unsafe"
)

type Loose struct {
	A bool  // 1 byte + 7 padding, so B starts at offset 8
	B int64 // 8
	C bool  // 1 byte + 7 padding: the size is rounded up to a multiple of 8
}

type Tight struct {
	B int64 // 8
	A bool  // 1
	C bool  // 1 + 6 padding
}

func main() {
	fmt.Println("Loose:", unsafe.Sizeof(Loose{})) // 24
	fmt.Println("Tight:", unsafe.Sizeof(Tight{})) // 16
}
