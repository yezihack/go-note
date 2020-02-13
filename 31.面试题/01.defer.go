package main

import "fmt"

func print_hello() {
	defer func() {
		fmt.Println("a")
	}()
	go func() {
		fmt.Println("goroutine")
	}()
	go func() {
		defer fmt.Println("C")
		fmt.Println("D")
	}()
	//defer func() {recover()}()
	//defer recover() //输出panic信息的
	//go func() {recover()}() //输出panic信息的
	defer func() {fmt.Println("b")}()
	panic("c")
}
/*
output:
输出b, a
*/
func main() {
	print_hello()
	//
}