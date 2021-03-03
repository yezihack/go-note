package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		for v := range ch {
			fmt.Println("val", v)
		}
		fmt.Println("end")
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	time.Sleep(time.Second)
}
