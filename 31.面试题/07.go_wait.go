package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i ++ {
		go func() {
			fmt.Println("i", i) //i 可能一直是最后一个数.
			wg.Done()
		}() //内置函数时直接使用局部变量，未进行参数传递
	}
	for i := 0; i < 10; i ++ {
		go func(i int) {
			fmt.Println("j", i)
			wg.Done()
		}(i) //进行了参数传递.
	}
	wg.Wait()
}
/*
struct{}类型值的表示法只有一个，即：struct{}{}。
并且，它占用的内存空间是0字节。
确切地说，这个值在整个 Go 程序中永远都只会存在一份。
虽然我们可以无数次地使用这个值字面量，但是用到的却都是同一个值。
*/