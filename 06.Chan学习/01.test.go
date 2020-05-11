package main

import "fmt"

func main() {
	ch := make(chan int,3)
	ch <- 1
	ch <- 1
	ch <- 1
	close(ch) //如果不关闭通道的话,使用For-range语法, 会阻塞操作的.
	show(ch)
}
func show(c chan<- int) {
	for x := range c {
		fmt.Println(x)
	}
}
