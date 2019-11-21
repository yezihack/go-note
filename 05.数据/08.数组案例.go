package main

import "fmt"

func main() {
	var s [5]string
	s[0] = "wang"
	s[1] = "wang"
	s[2] = "wang"
	s[3] = "wang"
	s[4] = "wang"
	fmt.Printf("ss:%#v, type:%T\n", s, s)       //还是一个数组
	fmt.Printf("ss:%#v, type:%T\n", s[:], s[:]) //数组已经转换成一个切片啦.
}
