package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	timeout4()
}
func timeout5() {
	ch := make(chan int, 10)
	go func() {
		time.Sleep(time.Second * 3)
		ch <- 10
	}()
	to := time.NewTimer(1e9)
	for {
		select {
		case ret := <-ch:
			fmt.Println("ret:", ret)
		case <-to.C:
			fmt.Println("timeout")
			to.Reset(1e9)
		}
	}
}

func timeout4() {
	c := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			c <- 1
			time.Sleep(time.Second * 2)
		}

		os.Exit(0)
	}()

	for {
		select {
		case n := <-c:
			println("n:", n)
		case <-timeAfter(time.Second):
		}
	}
}
func timeAfter(d time.Duration) chan int {
	q := make(chan int, 1)

	time.AfterFunc(d, func() {
		q <- 1
		println("run") 		// 重点在这里
	})

	return q
}
func timeout3() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i ++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()
	select {
	case c := <-ch:
		fmt.Println("ch-int", c)
	case <- timeAfter2(2e9):
		fmt.Println("timeout")
	}
}
func timeAfter2(d time.Duration) chan int {
	q := make(chan int, 1)
	time.AfterFunc(d, func() {
		q <- 1
		fmt.Println("run")
	})
	return q
}
func timeout2() {
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println(<-ch)
	case <-time.After(1e9):
		fmt.Println("timeout")
	}
}
func timeout1() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	ch := make(chan int)
	select { //借用select解决越时问题.
	case <-ch:
		fmt.Println("ch")
	case <-timeout:
		fmt.Println("timeout")
	}
}
