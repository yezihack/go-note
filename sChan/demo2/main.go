package main

import (
	"fmt"
	"time"
)

var c chan bool

func main() {
	c = make(chan bool)
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("expense")
			}
		}
	}()
	display()
	<-time.After(time.Millisecond * 500)
}

func display() {
	c <- true
}
