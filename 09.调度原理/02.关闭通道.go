package main

import "fmt"

//通道需要关闭.这是一个良好的习惯.否则可能会导致资源泄露.
func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for a := range c {
		fmt.Println(a)
	}
}
