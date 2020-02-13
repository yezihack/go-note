package main

import "fmt"

func main() {
	c := make(chan int, 2) // 一个容量为2的缓冲通道
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok) // 3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok) // 5 true
	fmt.Println(len(c), cap(c)) // 0 2
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
	fmt.Println(len(c), cap(c)) // 0 2
	close(c) // 此行将产生一个恐慌 向一个关闭的通道再次关闭则会panic
	c <- 7   // 如果上一行不存在，此行也将产生一个恐慌。//向一个关闭的通道发送操作 ,也会panic

}