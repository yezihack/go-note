package main

import (
	"fmt"
	"sync"
)
//close通道,使用ok判断通道是否关闭. 起到广播作用.

//生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i ++ {
			ch <- i
		}
		//发送完了,关闭通道,起到广播作用.
		close(ch)
		//ch<-1 //向关闭的通道,会出现panic
		wg.Done()
	}()
}
func dataConsumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			//这里利用检查通道是否关闭, 如果关闭则不再读取数据啦.
			if i, ok := <-ch; ok {
				fmt.Printf("i=%d\n", i)
			} else {
				break //跳出循环
			}
		}
		wg.Done()
	}()
}
func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, wg)
	wg.Add(1)
	dataConsumer(ch, wg)
	wg.Add(1)
	dataConsumer(ch, wg)
	wg.Add(1)
	dataConsumer(ch, wg)
	wg.Wait()
}
