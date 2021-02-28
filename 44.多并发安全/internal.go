package main

type GoSafe struct {
	valChan chan int // 管道传递的类型
	result  []int    // 装载结果
}
