package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup

func main() {
	wg = new(sync.WaitGroup)
	c := make(chan int)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(workId int, wgChild *sync.WaitGroup) {
			defer wgChild.Done()
			for v := range c {
				fmt.Printf("work: %d, value:%d\n", workId, v)
				time.Sleep(time.Millisecond * 500)
			}
		}(i, wg)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c) //必须关闭通道
	wg.Wait()
}
