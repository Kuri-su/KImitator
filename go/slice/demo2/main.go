package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

func main() {
	path := []byte("AAAA/BBBBBB")
	fmt.Printf("raws: %p,len: %d, cap: %d, %v \n", path, len(path), cap(path), string(path))
	sepIndex := bytes.IndexByte(path, '/')
	fmt.Println("sepIndex", sepIndex)

	dir1 := path[:sepIndex:sepIndex]
	dir2 := path[sepIndex+1:]

	fmt.Println(unsafe.Pointer(&path[0]))
	fmt.Println(unsafe.Pointer(&dir1[0]))
	fmt.Printf("dir1: %p,len: %d, cap: %d, %v , %s \n", dir1, len(dir1), cap(dir1), dir1, string(dir1))
	fmt.Printf("dir2: %p,len: %d, cap: %d, %v , %s \n", dir2, len(dir2), cap(dir2), dir2, string(dir2))

	dir1 = append(dir1, "suffix"...)

	fmt.Println(unsafe.Pointer(&path[0]))
	fmt.Println(unsafe.Pointer(&dir1[0]))
	fmt.Println("raws =>", string(path))
	fmt.Printf("dir1: %p,len: %d, cap: %d, %v , %s \n", dir1, len(dir1), cap(dir1), dir1, string(dir1))
	fmt.Printf("dir2: %p,len: %d, cap: %d, %v , %s \n", dir2, len(dir2), cap(dir2), dir2, string(dir2))
}
