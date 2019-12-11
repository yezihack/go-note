package _8_Goroutine注意

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//count变量在多个协程中运行,是线程不安全的.
func TestCounter(t *testing.T) {
	count := 0
	for i := 0; i < 5000; i ++ {
		go func() {
			count ++
		}()
	}
	fmt.Println(count)
}
//需要保证count变量安全,则需要加上锁机制.
func TestCounterThreadSafe(t *testing.T) {
	count := 0
	var mut sync.Mutex
	for i := 0; i < 5000; i ++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count ++
		}()
	}
	//我们不知道5000个协程需要多长的时间,需要我们假设一个等待时间.
	time.Sleep(1 * time.Second)
	t.Log("counter=", count)

}
//使用waitGroup监控goroutine什么时候执行完.
//包括三个方法 add,  done, wait
func TestCounterWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i ++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter=%d", counter)
}