package main

import "fmt"

//进程. 独立单元,内存独立.稳定性强.系统调度
//线程. 线程与线程之间内存共享,存在线程安全问题,需要加锁.上下文切换内存消耗大,系统调度.共享内存不安全.
//协程. 逻辑并发

//同步: 即串行执行
//异步: 是并发执行.
func main() {
	go func() {
		fmt.Println("hello")
	}()
	fmt.Println("ok")
}
