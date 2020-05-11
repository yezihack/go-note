package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ChanIsClose() {
	var ch chan int
	ch = make(chan int, 1)
	ch <- 1
	close(ch)//如果未关闭的话, 会阻塞.因为接收协程队列为空.
	//除非缓冲未满状态或者 接收协程队列不为空,则不会出现阻塞.
	if val, ok := <-ch; ok {
		fmt.Println(val)
	}
	if val, ok := <-ch; ok {
		fmt.Println("ok", val)
	} else {
		fmt.Println("通道已送")
	}
}
//
func ChanIsCloseV2() {
	ch := make(chan int)
	go func() { //开一个接收的协程队列.不关闭通道也不会阻塞
		for i := 0;i < 2; i ++ {
			ch <- rand.Intn(100)
			time.Sleep(time.Second)
		}
	}()
	if val, ok := <-ch; ok {
		fmt.Println(val)
	}
	if val, ok := <-ch; ok {
		fmt.Println("ok", val)
	} else {
		fmt.Println("通道已送")
	}
	close(ch) //如果没有这一行,则再去接收数据的话,会出现阻塞deadlock
	if val, ok := <-ch; ok {
		fmt.Println("ok", val)
	} else {
		fmt.Println("通道已送")
	}
}
func main() {
	//rand.Seed(time.Now().UnixNano())
	//ChanIsCloseV2()
}
