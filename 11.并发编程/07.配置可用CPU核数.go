package main

import (
	"fmt"
	"runtime"
)

//https://studygolang.com/pkgdoc
//查看逻辑的内核数runtime.NumCPU()
func main() {

	fmt.Println("当前逻辑内核数", runtime.NumCPU())
	//设置当前协程可用的CPU逻辑核心数,并返回之前的逻辑数
	fmt.Println("之前的逻辑数:", runtime.GOMAXPROCS(4))
	fmt.Println("之前的逻辑数:", runtime.GOMAXPROCS(2))
	fmt.Println("之前的逻辑数:", runtime.GOMAXPROCS(1))
	fmt.Println("之前的逻辑数:", runtime.GOMAXPROCS(0))
}
