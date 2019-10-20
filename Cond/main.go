package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = new(sync.Mutex)
var cond = sync.NewCond(lock)
var number = 10

func main() {
	for i := 0; i < number; i++ {
		go func() {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			i++
			fmt.Printf("my is %d goroutine\n", i)
		}()
	}
	fmt.Println("start one sign")
	time.Sleep(time.Second)
	cond.Signal()

	fmt.Println("start second sign")
	cond.Signal()
	time.Sleep(time.Second)
	cond.Signal()
	fmt.Println("start all")
	time.Sleep(time.Second * 2)
	cond.Broadcast()

}
