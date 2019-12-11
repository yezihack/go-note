package main

import (
	"fmt"
	"time"
)

func isCancelled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}
func doCancel(ch chan struct{}) {
	ch <- struct{}{}
}
func doCancel2(ch chan struct{}) {
	close(ch)
}

func main() {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5;i ++{
		go func(i int, cancelCh  chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println("Cancel i=", i)
		}(i, cancelChan)
	}
	//doCancel(cancelChan)
	doCancel2(cancelChan)
	time.Sleep(time.Millisecond * 100)
}
