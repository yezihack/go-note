package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//新建一个关闭chan
	stop := make(chan struct{})
	//todo 业务
	Show(stop)

	//创建一个信号chan
	sign := make(chan os.Signal)
	//监听所有信号操作.
	signal.Notify(sign, os.Kill, os.Interrupt, syscall.SIGKILL)
	//释放信号
	<-sign
	//发起一个回收资源信号.
	stop <- struct{}{}
	fmt.Println("END")
}

//定时器实现
//Timer, Ticker的区别
//Ticker一但被定义, 每隔一段时间会自动触发.
//Timer 定义后固定只执行一次,使用Reset会再触发一次.
func Show(stop chan struct{}) {
	ticker := time.NewTicker(time.Second) //Ticker一但被定义, 每隔一段时间会自动触发.
	timer := time.NewTimer(time.Second)   //Timer 定义后固定只执行一次,使用Reset会再触发一次.
	//Ticker的实现
	go func(t *time.Ticker) {

		for {
			select {
			case <-t.C:
				fmt.Printf("ticker:%v\n", 11)
			}
		}
	}(ticker) //

	//Timer的实现
	go func(t *time.Timer) {
		for {
			select {
			case <-t.C:
				fmt.Println("timer:", 22)
				t.Reset(time.Second)
			}
		}
	}(timer)

	//开个协程,用于回收资源.
	go func(t *time.Ticker, t1 *time.Timer) {
		<-stop
		t.Stop()
		t1.Stop()
		fmt.Println("回收资源")
	}(ticker, timer) //以形参方式将外部的变量传递进来,避免变量逃逸

	//以下代码改为上面的写法.
	go func() {
		<-stop
		timer.Stop()
		ticker.Stop()
	}()
}

// 关于堆栈和指针上的语言力学 https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html
// go build -gcflags '-m -l' 逃逸分析
// go tool compile -m timer.go
