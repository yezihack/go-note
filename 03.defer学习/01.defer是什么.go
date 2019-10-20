package main

import "fmt"
//defer用来声明一个延迟函数, 先进后出.
func main() {
	CallDefer()
}
func CallDefer() {
	defer func() {
		fmt.Println("1")
	}()
	defer func() {
		fmt.Println("2")
	}()
	defer func() {
		fmt.Println("3")
	}()
	panic("error")
}
