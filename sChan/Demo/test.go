package main

import (
	"fmt"
	"time"
)

func main() {
	c, quit := make(chan int), make(chan int)
	go func() {
		i := 0
		for i < 10{
			i ++
			fmt.Println(<-c)
		}
		quit <- 99
	}()
	fib(c, quit)
}

func ChanFibonacci() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10 ; i ++ {
			fmt.Println(<- c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fib(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

func ChanClose() {
	go func() {
		time.Sleep(time.Hour)
	}()

	c := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i ++ {
			c <- i
		}
		close(c) //必须关掉才不会阻塞
	}()
	for i := range c {
		fmt.Println(i)
	}
}

func ChanAdd() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	c := make(chan int)
	go Add(s[:len(s)/2], c)
	go Add(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y)
}

func Add(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}