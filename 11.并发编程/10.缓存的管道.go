package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	fmt.Println(cap(ch))
}
