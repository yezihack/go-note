package main

import "fmt"

func Count(ch chan<- int) {
	fmt.Println("Count")
	ch <- 1
}
func GoCount() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i ++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		fmt.Println(<-ch)
	}
}
func chanSing() {
	ch := make(chan int, 2)
	_ = <-chan int(ch) //转换成一个单向的读取channel
	ch2 := chan<- int(ch) //转换成一个单向的写入channel
	ch2 <- 1
}
func main() {
	ch := make(chan int, 3)
	ch <- 2
	fmt.Println(<-ch)
	if val, ok := <-ch; ok {
		fmt.Println(val)
	}
	close(ch)
}
