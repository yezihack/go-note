package main

import (
	"fmt"
	"time"
)

//如果进行接收操作,而发送协程队列为空的话,则阻塞deadlock
func c1() {
	ch := make(chan int)
	fmt.Println(<-ch)
}
//如果进行发送操作, 而接收协程队列为空的话,则阻塞deadlock
func c2() {
	ch := make(chan int)
	ch<- 1
}
//如果进行接收操作,而发送协程队列不为空的话,正常接收操作.
func c3() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}
//如果进行发送操作, 而接收协程队列不为空的话,则正常发送操作.
func c4() {
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- 1
}
//如果缓冲队列为空的话, 接收操作的话,会阻塞, deadlock
func c5() {
	ch := make(chan int, 2)
	fmt.Println(<-ch)
}
//如果缓冲队列不为空而未满状态下, 正常接收操作.
func c6() {
	ch := make(chan int, 2)
	ch <- 1
	fmt.Println(<-ch)
}
//如果缓冲队列被取空数据状态下, 再接收数据的话,会阻塞deadlock
func c7() {
	ch := make(chan int, 2)
	ch <- 1
	//close(ch)如果想不阻塞则关闭, 无限接收数据.
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
//如果缓冲队列满,则发送协程队列不为空,则可以正常接收.
func c8() {
	ch := make(chan int, 2)
	go func() {
		ch <- 1
		ch <- 1
		ch <- 1
		ch <- 1
	}()
	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
func main() {
	c8()
}
