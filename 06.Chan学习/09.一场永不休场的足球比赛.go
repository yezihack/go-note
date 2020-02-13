package main

import (
	"fmt"
	"time"
)

func main() {
	ball := make(chan string)
	kickBall := func(payerName string) {
		for { //无限循环下去. c操作将永久阻塞.
			fmt.Print(<-ball, "传球", "\n")
			time.Sleep(time.Millisecond * 500)
			ball <- payerName
		}
	}
	go kickBall("刘备")
	go kickBall("张飞")
	go kickBall("关羽")
	go kickBall("诸葛亮")
	ball <- "裁判" //准备开球
	var c chan bool //零值nil通道
	<-c //永久阻塞在此
}
