package main

import (
	"fmt"
	"time"
)

func main() {
	MyGoroutine()
	time.Sleep(time.Second)
}
func MyGoroutine() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("不会导致程序中止运行")
		}
	}()
	//直接导致程序中止运行
	panic("出错啦")
}
