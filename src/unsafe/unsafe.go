package main

import (
	"fmt"
	"unsafe"
)

func main() {
	intsize64 := unsafe.Sizeof(int64(42))
	intsize32 := unsafe.Sizeof(int32(2))
	intsize := unsafe.Sizeof(int(2))

	fmt.Println("int64:", intsize64)
	fmt.Println("int32:", intsize32)
	fmt.Println("int:", intsize)
}
