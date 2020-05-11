package main

import (
	"fmt"
	"time"
)

// 什么是并行: 同一时间段,有多个任务同步执行.更加节省时间，效率更高
//go语言里使用goroutine实现

func show(i int) {
	fmt.Println("我是", i)
}
func main() {
	fmt.Println("开始并行操作")
	//以下是并发, 打印不是按顺序执行的.而是一个抢占式的执行.
	for i := 0; i < 10; i++ {
		go show(i)
	}
	time.Sleep(time.Millisecond * 100)
	fmt.Println("并行完成")
}
