package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func printOrder(i int) {
	fmt.Println(i)
}

//利用自旋锁实现并发按顺序打印.
func main()  {
	var count int32
	trigger := func(i int32, fn func()) {
		for {
			//如果与原子操作自增操作的值相等则执行函数
			if n := atomic.LoadInt32(&count); n == i {
				fn()
				atomic.AddInt32(&count, 1)
				break
			}
			//否则睡一会儿,再判断, 这是一种自旋操作.
			time.Sleep(time.Nanosecond)
		}
	}
	for i := int32(0); i < 10; i ++ {
		go func(i int32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	//time.Sleep(time.Second)
	trigger(10, func() { //当自旋的值到达10就退出.
		fmt.Println("Finish.")
	})
}
