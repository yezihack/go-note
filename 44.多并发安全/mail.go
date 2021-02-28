package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("并发")
	a1 := []int{1, 2}
	a2 := []int{3, 4}
	s1 := time.Now()
	//result1 := add(a1, a2)       // 普通相加 elapsed time:4.0256375s, result:20
	//result1 := addLock(a1, a2)   // 加锁版本 elapsed time:4.0295941s, result:20
	result1 := addGoSync(a1, a2) // 同步版本 elapsed time:1.00845s, result:20
	fmt.Printf("elapsed time:%v, result:%v\n", time.Since(s1), result1)
}

// elapsed time:4.0256375s, result:20
func add(arr1, arr2 []int) int {
	result := 0
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			time.Sleep(time.Second)
			fmt.Printf("v1:%d, v2:%d, %d+%d=%d\n", v1, v2, v1, v2, v1+v2)
			result += v1 + v2
		}
	}
	return result
}

// 加锁版本
// elapsed time:4.0295941s, result:20
func addLock(arr1, arr2 []int) int {
	var (
		result = 0
		lock   = sync.Mutex{}
		wg     = sync.WaitGroup{}
	)
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			wg.Add(1)
			go func(v1, v2 int) {
				defer wg.Done()
				lock.Lock()
				defer lock.Unlock()
				time.Sleep(time.Second)
				result += v1 + v2
				fmt.Printf("v1:%d, v2:%d, %d+%d=%d\n", v1, v2, v1, v2, v1+v2)
			}(v1, v2)
		}
	}
	wg.Wait()
	return result
}

// 采用同步通道实现
func addGoSync(arr1, arr2 []int) int {
	var (
		result  = 0                // 结果集
		retChan = make(chan int)   // 结果通道
		wg      = sync.WaitGroup{} // 同步等待组
	)
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			wg.Add(1) // +1
			go func(v1, v2 int) {
				defer wg.Done()         // 标记结束
				time.Sleep(time.Second) // 延时操作
				retChan <- v1 + v2      // 计算结果
				fmt.Printf("v1:%d, v2:%d, %d+%d=%d\n", v1, v2, v1, v2, v1+v2)
			}(v1, v2)
		}
	}
	// 收集通道里的数据
	go func() {
		for val := range retChan {
			result += val
		}
	}()
	wg.Wait() // 等待所有的 goroutine 执行完成。
	return result
}
