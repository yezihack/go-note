package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch1 := make(chan int, 1)
	go func() {
		for   {
			fmt.Println("1")
		}
	}()
	//ch <- 1
	//ch1 <- 2
	//close(ch)
	//close(ch1)
	select {
	case <-ch:
		fmt.Println("ch")
	case <-ch1:
		fmt.Println("ch1")
	//default:
	//	fmt.Println("default")
	}
}
