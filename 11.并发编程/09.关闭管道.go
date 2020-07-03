package main

import (
	"fmt"
	"time"
)

// 关闭通道原则: 由发送方关闭通道

// 管着被关闭,无法再写入数据,但可以读出.
// 向关闭的管道里写,报: panic: send on closed channel

// 不能关闭一个未初使化的管道 nil
// 不能向已关闭的管道写数据
// 不能重复关闭管道.
// 向一个已关闭的管道里读数据不会报错, 会读取零值.

// close关闭会通知读取chan的所有子协程.
// close具有广播能力.

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("开始读取")
		for x := range ch {
			fmt.Println("读取", x)
		}
		//如果没有关闭管道, 这一行永远不会被执行,因为一直阻塞在range那一行.21行
		fmt.Println("子协程结束 ")
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	//close关闭会通知读取chan的所有子协程.
	close(ch)
	for {
		time.Sleep(time.Second)
	}
}

//使用第二个参数, 判断管道里是否有值.
func main08() {
	ch := make(chan int)
	go func() {
		for {
			//管道关闭后,会读取零值.
			x, ok := <-ch
			if ok {
				fmt.Println(x)
			} else {
				fmt.Println("chan is close")
			}
		}
	}()
	ch <- 1
	ch <- 2
	close(ch)
	for {
		time.Sleep(time.Second)
	}
}
func main07() {
	ch := make(chan int)
	go func() {
		for {
			//管道关闭后,会读取零值.
			fmt.Println(<-ch)
		}
	}()
	ch <- 1
	ch <- 2
	close(ch)
	for {
		time.Sleep(time.Second)
	}
}

//向关闭的管道里写,报: panic: send on closed channel
func main06() {
	ch := make(chan int)
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	ch <- 1
	ch <- 2
	close(ch) //管道关闭后不能再写了, 会报 panic: send on closed channel
	ch <- 3
	for {
		time.Sleep(time.Second)
	}
}
