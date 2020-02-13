package main

import (
	"fmt"
	"time"
)

func s() {
	for i := 0; i < 5; i ++ {
		defer func(i int) {
			fmt.Println(i)
			go func(i int) {
				fmt.Println("i", i)
			}(i)
		}(i)
	}
	time.Sleep(time.Second * 2)
}
func main() {
	//startAt := time.Now()
	//defer fmt.Println(time.Since(startAt))
	//time.Sleep(time.Second)
	//输出os, 不符合我们预期, 因为defer是值传递.会立即对函数中引用的外部参数进行拷贝.
	//Since不是在退出之前计算, 而是在defer关键字调用时计算的.最终导致代码输出0s. 需要
	//解决这个问题,我们传入一个匿名函数. defer此时拷贝的是函数指针
	startAt := time.Now()
	defer func() {
		fmt.Println(time.Since(startAt))
	}()
	time.Sleep(time.Second)
	//input 1.0008488s. ok
}
