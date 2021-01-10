package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := make([]int, 32)
	b := a[0:16]
	a = append(a, 1)
	a[2] = 42
	fmt.Println(unsafe.Pointer(&a[0]))
	fmt.Println(unsafe.Pointer(&b[0]))

	fmt.Printf("%p,len: %d, cap: %d, %v \n", a, len(a), cap(a), a)
	fmt.Printf("%p,len: %d, cap: %d, %v \n", b, len(b), cap(b), b)
}
