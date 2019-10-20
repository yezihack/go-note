package main

import (
	"fmt"
)

//数组是值传递
func transValue(arr [5]int) {
	arr[0] = 100
	arr[1] = 300
	arr[2] = 400
	fmt.Println("值传递,函数内修改不会影响外部数组, arr: ", arr)
}
func transSlice(s []int) {
	s[0] = 99
	s[1] = 88
	s[2] = 77
	fmt.Println("引用传递,函数内的修改会影响外部切片, slice", s)
}

func main() {
	arr := [5]int{1,2,3, 4,5}
	s := []int{1,2,3,4,5}
	transValue(arr)
	transSlice(s)
	fmt.Println("值传递, 外部array", arr)
	fmt.Println("引用传递, 外部slice", s)
}
