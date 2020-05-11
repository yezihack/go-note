package _0_基础

import (
	"fmt"
	"math"
	"os"
	"sync"
	"testing"
	"time"
)

/*
以通信方式共享内存
调度器不能保证多个Goroutine执行次序, 且进程退出时不会等待它们结束.

*/
func GSum(id int) {
	var x int64
	for i := 0; i < math.MaxInt32; i ++ {
		x += int64(i)
	}
	println(id, x)
}
func TestGSum(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i ++ {
		go func(id int) {
			defer wg.Done()
			GSum(id)
		}(i)
	}
	wg.Wait()
}

func TestChan(t *testing.T) {
	data := make(chan int) //数据交换队列
	exit := make(chan bool)//退出通知
	go func() {
		for d := range data {//从队列迭代接收数据,直到close
			fmt.Println(d)
		}
		fmt.Println("recv over.")
		exit <- true //发出退出通知
	}()
	data <- 1
	data <- 2
	data <- 3
	data <- 4
	close(data)//关闭通道
	fmt.Println("send over.")
	<-exit //等待退出通知
}
//异步方式通过判断缓冲区来决定是否阻塞。如果缓冲区已满，发送被阻塞；缓冲区为空，接收被阻塞。
//带缓冲的chan, 性能更好, 避免大对象拷贝.
//通常情况下，异步 channel 可减少排队阻塞，具备更高的效率。但应该考虑使用指针规
// 避大对象拷⻉贝，将多个元素打包，减小缓冲区大小等。
func TestCacheChan(t *testing.T) {
	data := make(chan int, 3)
	exit := make(chan bool)
	data <- 1
	data <- 2
	data <- 3
	go func() {
		for d := range data {
			fmt.Println("d=", d)
		}
		exit <- true
	}()
	data <- 4
	data <- 5
	data <- 6
	close(data)
	<- exit
}
func TestChanSingle(t *testing.T) {
	c := make(chan int, 3)
	c <- 3
	fmt.Println(len(c), cap(c))
	var send chan<- int = c //只有发送能力,无接收能力
	var recv <-chan int = c //只有接收能力,无发送能力
	send <- 1
	fmt.Println(<-recv)
	fmt.Println(<-recv)
	//不能将单向的channel转换成普通的channel
	//d := (chan int)(send)
}
//如果需要同时处理多个 channel，可使⽤用 select 语句。它随机选择⼀一个可⽤用 channel 做
//收发操作，或执⾏行 default case。
func TestChanSelect(t *testing.T) {
	a, b := make(chan int, 3), make(chan int)
	go func() {
		v, ok, s := 0, false, ""
		for {
			select {/// 随机选择可⽤用 channel，接收数据
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}
			if ok {
				fmt.Println("s=", s, "v=", v)
			} else {
				os.Exit(0)
			}
		}
	}()
	//在循环中使⽤用 select default case 需要⼩小⼼心，避免形成洪⽔水。
	for i := 0; i < 5; i ++ {
		select {//// 随机选择可⽤用 channel，发送数据
		case a <- i:
		case b <- i:
		}
	}
	close(a)
	select {}
}
func TestTimeout(t *testing.T) {
	w := make(chan bool)
	c := make(chan int, 2)
	go func() {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout.")
		}
		w <- true
	}()
	c <- 1// 注释掉，引发 timeout。
	<-w
}