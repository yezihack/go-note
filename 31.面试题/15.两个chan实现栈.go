package main

import (
	"fmt"
	"time"
)

type Stack struct {
	ch1 chan int
	ch2 chan int
}
//两个chan如何实现一个栈的功能.
//先将数据发送给chan1, 然后再存储在切片里. 倒序转到chan2, 然后输出.

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	arr := make([]int, 0)
	for val := range ch1 {
		arr = append(arr, val)
	}
	fmt.Println(arr)
	for i := len(arr)-1; i >= 0; i -- {
		ch2 <- arr[i]
	}
	close(ch2)
	for val := range ch2 {
		fmt.Println(val)
	}
	time.Sleep(time.Millisecond*100)
}