package main

import (
	"fmt"
	"unsafe"
)

//空struct{}节约空间,不点宽度, 也就是说占用0字节内存空间
func main() {
	var i int64
	var array [3]int
	var slice []int
	var s struct{}
	fmt.Printf("int占用宽度:%dbit\n", unsafe.Sizeof(i))
	fmt.Printf("array占用宽度:%dbit\n", unsafe.Sizeof(array))
	fmt.Printf("slice占用宽度:%dbit\n", unsafe.Sizeof(slice))
	fmt.Printf("struct{}占用宽度:%dbit\n", unsafe.Sizeof(s))
	//两个空结构地址相等
	var a, b struct{}
	fmt.Printf("a的地址:%p, b的地址:%p\n", &a, &b)
}
