package main

import (
	"fmt"
	"time"
)

//一个非缓冲通道实现的请求/响应的例子
func main() {
	c := make(chan int)
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		//<-ch 编译不通过, 因为ch是单通道,只能发送操作.不能操作接收
		ch <- x * x
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch
		//ch <- 1 编译不通过, 因为ch是单通道, 只能接收操作.不能操作发送
		fmt.Println("n=", n)
		time.Sleep(time.Second)
		done <- struct{}{} //通知我已经接收完毕
	}(c)
	<-done
	fmt.Println("bye")
}