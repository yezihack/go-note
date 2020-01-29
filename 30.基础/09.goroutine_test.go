package _0_基础

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

//调⽤用 runtime.Goexit 将⽴立即终⽌止当前 goroutine 执⾏行，调度器确保所有已注册 defer
//延迟调⽤用被执⾏行

func TestGoroutine(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer println("A.defer")
		go func() {
			defer println("B.defer")
			runtime.Goexit()//runtime.Goexit 将⽴立即终⽌止当前 goroutine 执⾏行, 已注册的defer依然执行.
			println("B")
		}()
		println("A")
	}()
	wg.Wait()
	/*
A
B.defer
A.defer
	*/
}

//Gosched 让出底层线程，将当前 goroutine 暂停，放回队列等
//待下次被调度执⾏行

func TestGosched(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 6; i ++ {
			fmt.Println("i=", i)
			if i == 2 {
				runtime.Gosched() //让出底层线程, 将当前的goroutine暂停, 放回队列下次执行
			}
		}
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Hello Gosched")
	}()
	wg.Wait()
}