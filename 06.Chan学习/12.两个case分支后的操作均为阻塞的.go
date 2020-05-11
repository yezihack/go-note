package main

import "fmt"

func main() {
	var c chan struct{} // nil
	select {
	case <-c:             // 阻塞操作
	case c <- struct{}{}: // 阻塞操作
	default:
		fmt.Println("Go here.")
	}

}