package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func print1() {
	sign := make(chan struct{})
	for i := 0; i < 10; i ++ {
		go func(i int) {
			fmt.Println("i", i)
			sign <- struct{}{}
		}(i)
	}
	for i := 0; i < 10; i ++ {
		<-sign
	}
}
func autoSelf() {
	var count uint32
	trigger := func(i uint32, fn func()) {
		for { //自旋锁实现
			if n := atomic.LoadUint32(&count); n == i {
				fmt.Printf("n:%d, i:%d, count:%d\n", n, i, count)
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond) //睡1纳秒
		}
	}
	for i := uint32(0); i < 10; i ++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println("i", i)
			}
			trigger(i, fn)
		}(i)
	}
	trigger(20, func() {
	})
}
func print3(){
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println("i", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
func main() {
	//print1()
	//autoSelf() //自旋实现
	//print3()

}
