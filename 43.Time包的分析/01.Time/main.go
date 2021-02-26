package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	// 我们首先利用 time.NewTimer 构造了一个 2 秒的定时器，同时使用 <-timer.C 阻塞等待定时器的触发。
	<-time.NewTimer(time.Second).C

	fmt.Println(time.Now())
}
