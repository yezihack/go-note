package main

import (
	"fmt"
	"time"
)

//chan 分无缓冲channel,也叫同步channel
//缓冲channel, 异步channel
//一种先进先出的策略.
func main() {
	//1.同步channel,读与写必须并发操作.否则会deadlock
	c := make(chan int)
	go func() {
		fmt.Println("正在等待中")
		time.Sleep(time.Second * 3)
		fmt.Println("向通道里赋值")
		c <- 10
	}()
	fmt.Println("一直阻塞, 直到有数据过来")
	<-c

	//2. 缓冲channel, 写入的数据未满缓冲的大小, 可以异步读. 如果超过缓冲大小,必须采用同步写,否则deadlock
	fmt.Println("以下是带缓冲通道功能的channel")
	cache := make(chan int, 1)
	cache <- 10
	go func() {
		cache <- 22
	}()
	fmt.Println(<-cache)
}
