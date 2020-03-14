package main

import (
	"fmt"
	"time"
)

//异步回调:异步完成后再调用一个函数
func sayHello() {
	fmt.Println("程序结束")
}
func main() {

	go func(f func()) {
		for i := 0; i < 10; i++ {
			fmt.Println("i=", i)
		}
		f() //这是执行在goroutine里
	}(sayHello)
	time.Sleep(time.Second * 1)
}

