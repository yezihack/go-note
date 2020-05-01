package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"
)

func main() {
	var tag int
	flag.IntVar(&tag, "tag", 1, "tag 1 or 2")
	flag.Parse()
	switch tag {
	case 1:
		FromChanReceiveValue()
	case 2:
		ToChanSendValue()
	case 3:
		Init()
	}

}
// 从一个通道接收一个值来实现单对单通知.
func FromChanReceiveValue() {
	fmt.Println("从一个通道接收一个值来实现单对单通知")
	// 创建一个非缓冲的chan, 也就是说必须收发需要同步, 否则阻塞中.
	done := make(chan struct{})
	// 创建一个goroutine
	go func() {
		fmt.Println("A")
		t := time.NewTicker(time.Millisecond * 10)
		i := 0
		for {
			i ++
			fmt.Print(".")
			if i > 100 {
				break
			}
			<-t.C
		}
		<-done // 对chan进行接收操作.解除"代码(2)处"的阻塞
	}()
	done <- struct{}{} // 代码(2)处. 此处发送一个chan, 因为是非缓冲的chan, 必须等待接收方收到,否则一直等待中.
	fmt.Println("over!")
}
// 向一个通道发送一个值来实现单对单通知
func ToChanSendValue() {
	fmt.Println("向一个通道发送一个值来实现单对单通知")
	rand.Seed(time.Now().UnixNano())
	values := make([]byte, 32 * 1024 * 1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	done := make(chan struct{}) //做信号通知用的
	go func() {
		fmt.Println("排序中....")
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // 向chan发送操作. 解除(2)处的阻塞.
	}()
	<-done  // (2) 对chan进行接收操作. 非缓冲的chan, 必须等待发送方完成, 否则一直阻塞.
	//fmt.Println(values)
	fmt.Println(values[0], values[len(values)-1])
}

type Conf struct {
	Close chan struct{}
}

func Init() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	conf := Conf{}
	conf.Close = make(chan struct{})
	CloseDB(conf)
	go func() {
		t := time.NewTimer(time.Second * 3)
		select {
		case <-t.C:
			fmt.Println("向ch发送一个信息")
			conf.Close <- struct{}{}
		}
		wg.Done()
	}()
	wg.Wait()
	time.Sleep(time.Second)
}

var Done chan struct{}
func CloseDB(cfg Conf) {
	go func() {
		fmt.Println("=============等待收到发送操作.1")
		<-cfg.Close
		Done <- struct{}{}
		fmt.Println("=============等待收到发送操作.2")
	}()
	go func() {
		<-Done
		fmt.Println("--------------------------Done")
	}()
}

func Fib(c chan int, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}