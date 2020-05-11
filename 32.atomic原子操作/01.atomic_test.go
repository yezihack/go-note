package _2_atomic

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	counter int64
	wg sync.WaitGroup
	mutex sync.Mutex
)
func IncCounter(id int) {
	defer wg.Done()
	//这个函数会同步整型值的加法，
	//方法是强制同一时刻只能有一个 goroutine 运行并完成这个加法操作。
	//// 安全地对 counter 加 1
	atomic.AddInt64(&counter, 1)
	//// 当前 goroutine 从线程退出，并放回到队列
	runtime.Gosched()
	//counter ++
}
func IncCounter2(id int) {
	defer wg.Done()
	for count := 0; count < 2; count ++ {
		mutex.Lock()
		{
			value := counter
			// 当前 goroutine 从线程退出，并放回到队列
			runtime.Gosched()
			value ++
			counter = value
		}
		mutex.Unlock()
	}
}
func TestIncCounter(t *testing.T) {
	wg.Add(2)
	go IncCounter(1)
	go IncCounter(2)
	wg.Wait()
	fmt.Println("finish Counter: ", counter)
}
func TestAddInt64(t *testing.T) {
	var count int64
	atomic.AddInt64(&count, 1)
	atomic.AddInt64(&count, 1)
	atomic.AddInt64(&count, 1)
	fmt.Println(count)
}
var (
	shutdown int64
)
func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Println("Work:", name)
		time.Sleep(time.Millisecond * 250)
		//要停止工作吗
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Println("Work stop: ", name)
			break
		}
	}
}
func TestDoWork(t *testing.T) {
	wg.Add(2)
	go doWork("A")
	go doWork("B")
	//
	time.Sleep(time.Second)
	fmt.Println("shutdown now")
	atomic.StoreInt64(&shutdown, 1)
	wg.Wait()
}
func TestNum(t *testing.T) {
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
}
func TestFanMan(t *testing.T) {
	a := 6
	fmt.Println(uint32(int32(a)))
	fmt.Println(^uint32(-a-1))
}
