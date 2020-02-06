package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	//一秒后关闭通道
	time.AfterFunc(time.Millisecond * 100, func() {
		close(ch)
	})
	go func() {
		time.Sleep(time.Millisecond * 150)
		ch <- 100
	}()
	select {
	case val := <-ch:
		fmt.Println("value:", val)
	case _, ok := <-ch: //如果超出1秒则会关闭
		if !ok {
			fmt.Println("通道已经关闭")
			break
		}
		fmt.Println("end")
	}
}
func sum(ch chan int) {
	ch <- 100
}
