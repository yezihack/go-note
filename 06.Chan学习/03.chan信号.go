package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	go func() {
		fmt.Println("11")
		stop <- struct{}{}
		<-stop
		fmt.Println("22")
	}()
	go func() {
		<-stop
		fmt.Println("33")
	}()
	time.Sleep(time.Second)
}
