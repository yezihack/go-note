package main

import (
	"fmt"
	"runtime"
)
/*
会出现异常. 因为select的case选择是一种伪随机
*/
func main() {
	runtime.GOMAXPROCS(1)
	for {
		int_chan := make(chan int, 1)
		str_chan := make(chan string, 1)
		int_chan <- 1
		str_chan <- "aa"
		select {
		case val := <-int_chan:
			fmt.Println("int", val)
		case ret := <-str_chan:
			//fmt.Println("str", ret)
			panic("ret: "+ret)
		}
	}
}
