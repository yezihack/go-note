package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		//fmt.Println(<-ch)
		for {
			select {
			case v, ok := <-ch:
				if ok {
					fmt.Println("v", v)
				}
			default:
				goto END
			}
		END:
			break
		}
		//for v := range ch {
		//	fmt.Println("v", v)
		//}
	}()
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1

	close(ch)
	wg.Wait()
	fmt.Println("end")
}

func main() {

	fmt.Println("并发")
	a1 := []int{1, 2}
	a2 := []int{3, 4, 5}
	total := 0
	for range a1 {
		for range a2 {
			total++
		}
	}
	fmt.Println(total)
	s1 := time.Now()
	//result1 := add(a1, a2)       // 普通相加 elapsed time:4.0256375s, result:20
	//result1 := addLock(a1, a2)   // 加锁版本 elapsed time:4.0295941s, result:20
	result1 := addGoSync(a1, a2) // 同步版本 elapsed time:1.00845s, result:20
	//result1 := addGoJob(a1, a2) // 利用调度器实现多并发安全
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
		result  = 0                                   // 结果集
		retChan = make(chan int, len(arr1)*len(arr2)) // 结果通道
		wg      = sync.WaitGroup{}                    // 同步等待组
		done    = make(chan struct{})
	)
	go func() {
		// 收集通道里的数据
		// 如果 retChan 没有数据的话，一直会阻塞中。直到有数据过来才取数据。当 chan 关闭后才结束阻塞。
		//for {
		//	b := false
		//	select {
		//	case val, ok := <-retChan:
		//		if ok {
		//			result += val
		//		} else {
		//			b = true
		//			break
		//		}
		//	}
		//	if b {
		//		break
		//	}
		//}
		for val := range retChan {
			result += val
		}
		done <- struct{}{}
	}()
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
	wg.Wait() // 等待所有的 goroutine 执行完成。
	close(retChan)
	<-done
	return result
}

// 利用调度器实现多并发安全。
func addGoJob(arr1, arr2 []int) int {
	var (
		done = make(chan struct{}) // 结束标记
		size = len(arr1) * len(arr2)
		s    = NewScheduleJob(size, func() {
			done <- struct{}{}
		})
		wg = sync.WaitGroup{}
	)
	wg.Add(size)
	fmt.Println("size", size)
	for _, v1 := range arr1 {
		for _, v2 := range arr2 {
			go func(v1, v2 int) {
				defer wg.Done()         // 标记结束
				time.Sleep(time.Second) // 延时操作
				s.Add(v1 + v2)          // 计算结果
				fmt.Printf("v1:%d, v2:%d, %d+%d=%d\n", v1, v2, v1, v2, v1+v2)
			}(v1, v2)
		}
	}
	wg.Wait()
	s.Close()
	<-done
	result := s.Result()
	x := 0
	for _, v := range result {
		x += v
	}
	return x
}

// 调度任务
type ScheduleJob struct {
	ch   chan int // 通道传递数据类型
	data []int    // 装载数据
}

func NewScheduleJob(size int, done func()) *ScheduleJob {
	s := ScheduleJob{
		ch:   make(chan int, size),
		data: make([]int, 0),
	}
	go func() {
		s.collect()
		done()
	}()
	return &s
}

// 收集数据
func (s *ScheduleJob) collect() {
	for v := range s.ch {
		s.data = append(s.data, v)
	}
}
func (s *ScheduleJob) Result() []int {
	return s.data
}

// 关闭通道
func (s *ScheduleJob) Close() {
	close(s.ch)
}

// 添加数据
func (s *ScheduleJob) Add(v int) {
	s.ch <- v
}
