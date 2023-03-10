package main

import (
	"fmt"
	"unsafe"
)

type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	c int32
	b int16
}

func main() {
	fmt.Println(unsafe.Sizeof(demo1{}))  // 8
	fmt.Println(unsafe.Sizeof(demo2{}))  // 12
	fmt.Println(unsafe.Alignof(demo1{})) // 4
	fmt.Println(unsafe.Alignof(demo2{})) // 4
}
