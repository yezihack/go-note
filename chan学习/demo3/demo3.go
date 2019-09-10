package demo3

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg *sync.WaitGroup

func demoChan() {
	wg = new(sync.WaitGroup)
	c := make(chan int, 10)
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
	//close(c)
	wg.Wait()
}

func ChanClick() {
	type req struct {
		pId int
		val int
	}
	data := make(chan req, 10)
	//消费chan
	for i := 0; i < 3; i++ {
		go func(workId int) {
			for v := range data {
				fmt.Printf("workId: %d, from_workId: %d, value: %d\n", workId, v.pId, v.val)
			}
		}(i)
	}
	//生产chan
	for j := 0; j < 10; j++ {
		go func(workId int, reqChan chan req) {
			for {
				t := time.NewTicker(time.Millisecond * 500)
				select {
				case <-t.C:
					reqChan <- req{
						pId: workId,
						val: rand.Intn(100),
					}
				}
			}
		}(j, data)
	}
	select {}
}
