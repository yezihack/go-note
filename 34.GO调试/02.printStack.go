package main

import (
	"fmt"
	"runtime/debug"
)

func s() {
	fmt.Println("A")
	s1()
}
func s1() {
	fmt.Println("B")
	s2()
}
func s2() {
	fmt.Println("C")
	fmt.Printf("%s", debug.Stack()) //返回栈信息
	debug.PrintStack() //直接输出堆栈信息.
}
func main() {
	s()
}
