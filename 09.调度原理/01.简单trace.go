package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

//1. go run 01.go
//2. go tool trace trace.out
func main() {
	// 创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	c := make(chan int, 1)
	c <- 1

	// main
	fmt.Println("Hello trace")
}
