package main

import (
	"fmt"
	"runtime"
	"time"
)

//如果开了协程,使用chan通信,chan只取走一个,后面的无法取走,则其它协程会一直阻塞
//导致协程泄露.

func runTask(i int) string {
	time.Sleep(time.Millisecond * 100)
	return fmt.Sprintf("task id:%d\n", i)
}
func FirstResponse() string {
	numOfRunner := 10
	//ch := make(chan string)//如果无缓冲, 以下开10个协程运行,而只返回一个chan, 还有9个一直等待,因为没有取走数据,这样做会出现协程泄露.
	ch := make(chan string, numOfRunner) //加入缓冲channel,则以下协程不会阻塞.
	for i := 0; i < numOfRunner; i ++{
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}
func main() {
	fmt.Println("before", runtime.NumGoroutine())
	fmt.Println(FirstResponse())
	time.Sleep(time.Millisecond * 500)
	fmt.Println("after", runtime.NumGoroutine())
}