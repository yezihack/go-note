package main

import (
	"fmt"
	"time"
)
//defer 延时执行函数, 一般用于收尾工作.
func panic1() {
	defer fmt.Println("A")
	go func() {
		defer fmt.Println("B")
		panic("panic")
	}()
	time.Sleep(time.Millisecond * 1000)
}
//recover只能在goroutine里使用.其它地方使用不会起作用.
//只要在goroutine里执行才有用
func panic_recover() {
	defer fmt.Println("A")
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	panic("panic")
}
func panic_recover_v2() {
	/***************能起到恢复作用****************/
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()
	/***************起不到恢复作用****************/
	//go func() {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			fmt.Println(err)
	//		}
	//	}()
	//}()
	/***************起不到恢复作用****************/
	//defer fmt.Println("in main")
	//panic("unknown err")
	//if err := recover(); err != nil {
	//	fmt.Println(err)
	//}
	/***************起不到恢复作用****************/
	/*
	多次调用 panic 也不会影响 defer 函数的正常执行
output:
in main
panic: panic C
	panic: panic B
	panic: panic A
	*/
	defer println("in main")
	defer func() {
		defer func() {
			panic("panic A")
		}()
		panic("panic B")
	}()
	panic("panic C")
}
func panic_v3() {
	defer panic(nil)
	defer panic("panic A")
	defer panic("panic B")
}
func main() {
	//panic1()
	//panic_recover()
	//panic_recover_v2()
	panic_v3()
}
