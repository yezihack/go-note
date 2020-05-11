package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)//把信道关闭掉,否则deadlock,不关闭的话,信处还是有能力发送的.的range处会阻塞.
	for c := range ch {
		fmt.Println(c)
	}
}
