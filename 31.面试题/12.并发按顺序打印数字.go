package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//原子操作的变量.
var Count int32
//实现一个自旋锁操作.
func SpinLock(i int32, fn func()) {
	for { //一个死循环.
		//查看原子操作的值.如果相等则执行函数
		if n := atomic.LoadInt32(&Count); n == i {
			fn()
			atomic.AddInt32(&Count, 1) //然后原子操作自增.
			break //一定要跳出循环.
		}
		time.Sleep(time.Nanosecond)
	}
}
func main() {
	for i := int32(0); i < 10; i ++ {
		go func(i int32) {
			fn := func() {
				fmt.Println(i)
			}
			SpinLock(i, fn)
		}(i)
	}
	SpinLock(10, func() {})
}