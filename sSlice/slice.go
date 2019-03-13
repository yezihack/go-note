package sSlice

import (
	"fmt"
	"unsafe"
)

/******************************
*len 获取slice的长度
cap获取slice的最大容量
append向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
copy函数从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
 ******************************/
func StudySlice() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //声明一个切片,并初使一些值
	fmt.Printf("a切片: %v, 长度:%d, 切片cap:%d, 地址: %p\n", a, len(a), cap(a), a)
	a1 := a[1:4]
	fmt.Printf("a1切片: %v, 长度:%d, 切片cap:%d, 地址: %p\n", a1, len(a1), cap(a1), a1)
	a1 = append(a1, 5, 6, 7, 8, 9, 10, 5, 6, 7, 8, 9, 10, 5, 6, 7, 8, 9, 10)
	fmt.Printf("a1 append后切片: %v, 长度:%d, 切片cap:%d, 地址: %p\n", a1, len(a1), cap(a1), a1)
	a2 := make([]int, 0)
	a2 = a1[:]
	fmt.Printf("a1地址:%p, a2地址:%p\n", a1, a2)
	prt := unsafe.Pointer(&a2[0])
	fmt.Printf("a2[0]:%p\n", prt)
}

func ForArray() [1024]int {
	var a [1024]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	return a
}
func ForSlice() []int {
	a := make([]int, 1024)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	return a
}
