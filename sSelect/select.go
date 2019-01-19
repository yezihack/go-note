package sSelect

import (
	"fmt"
	"time"
	"runtime"
)
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

//fatal error: all goroutines are asleep - deadlock!
//如果协程里没有数据则会报错.
func StudySelect() {
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
	}()
	select {
	case <-ch1:
		fmt.Println("我是ch1")
	case <-ch2:
		fmt.Println("我是ch2")
	}
}

//select一直会等待数据接受
func StudySelect2() {
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		ch1 <- true
	}()
	select {
	case <-ch1:
		fmt.Println("我是ch1")
	case <-ch2:
		fmt.Println("我是ch2")
	}
}

func StudySelect3() {
	ch1 := make(chan int, 1)
	ch1 <- 1
	select {
	case ch1 <- 2:
	default:
		fmt.Println("chan is full...")
	}
}
