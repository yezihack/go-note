package main

import (
	"fmt"
	"time"
)

func AsyncService() chan string{
	retCh := make(chan string)
	go func() {
		time.Sleep(time.Millisecond * 500)
		retCh <- "abc"
	}()
	return retCh
}
func main() {
	ret := AsyncService()
	for {
		select {
		case str := <-ret:
			fmt.Println("result", str)
			break
		case <-time.After(time.Millisecond * 50):
			fmt.Println("wait...")
		}
	}
}
