/*
1. 管道: 是一个引用.
2. 管道:make 创建了一个底层数据结构的引用,当赋值或参数传递时,只是拷贝一个channel引用
指向相同的channel对象.
3. 向管道里写数据 ch <- 1
4. 读取管道里的数据: value := <- ch
5. 管道分为: 无缓冲管道, 有缓冲的管道.
*/
package main

import (
	"fmt"
	"time"
)

//主协程写3个, 子协程读5个, 子协程会永远阻塞
func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			//如果读空了, 没人写, 则读出阻塞
			x := <-ch //主协程只写3次, 没有数据啦, 一直阻塞在读出阶段.
			fmt.Println("读出", x)
		}
	}()
	for i := 0; i < 3; i++ {
		//缓存已满,没有人读,则写入阻塞
		ch <- i
		fmt.Println("写入", i)
		time.Sleep(time.Millisecond * 100)
	}
	for {
		time.Sleep(time.Second)
	}
}

//正常的读取. 主协程写, 子协程读.
func main05() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			//如果读空了, 没人写, 则读出阻塞
			x := <-ch
			fmt.Println("读出x=", x)
		}
	}()
	for i := 0; i < 5; i++ {
		//缓存已满,没有人读,则写入阻塞
		ch <- i
		fmt.Println("写入i=", i)
		time.Sleep(time.Millisecond * 100)
	}

}

func main04() {
	//无缓冲能力的管道,必须有一个协程去读, 一个写,同时进行.否则死锁
	//ch := make(chan int)
	//ch <- 1

	//带缓存能力的管道,如果超出缓存能力,又无其它协程写的话会出现死锁.
	ch := make(chan int, 1)
	ch <- 3
	fmt.Println(<-ch)
	time.Sleep(time.Millisecond * 1000)
}
