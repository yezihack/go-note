package main

import (
	"fmt"
	"strconv"
	"sync"
)

type UserAges struct {
	ages map[string]int
	//sync.Mutex
	sync.RWMutex
}
func (u *UserAges) Set(key string, value int) {
	u.Lock()
	defer u.Unlock()
	u.ages[key] = value
}
func (u *UserAges) Get(key string) int {
	u.RLock()
	defer u.RUnlock()
	if val, ok := u.ages[key]; ok {
		return val
	}
	return -1
}
//以下代码会不会出错, 出错如何修改.
//func main() {
//	//初使化操作.
//	u := UserAges{ages: map[string]int{}}
//	//添加操作.
//	for i := 0; i < 1000;i ++ {
//		go func(i int) {
//			u.Set("u"+strconv.Itoa(i), i)
//		}(i)
//	}
//	//获取操作
//	for i := 0; i < 1000;i ++ {
//		go func(i int) {
//			fmt.Println(u.Get("u"+strconv.Itoa(i)))
//		}(i)
//	}
//	fmt.Println("finish")
//}
//以上代码会出现goroutine错误.我们加入sync.WaitGroup
//func main() {
//	//初使化操作.
//	u := UserAges{ages: map[string]int{}}
//	wg := sync.WaitGroup{}
//	//添加操作.
//	for i := 0; i < 1000;i ++ {
//		wg.Add(1)
//		go func(i int) {
//			u.Set("u"+strconv.Itoa(i), i)
//			wg.Done()
//		}(i)
//	}
//	//获取操作
//	for i := 0; i < 1000;i ++ {
//		wg.Add(1)
//		go func(i int) {
//			fmt.Println(u.Get("u"+strconv.Itoa(i)))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	fmt.Println("finish")
//}
//以上代码还是会出错. 因为使用了goroutine操作, 读取时会出现并发安全问题.
// 我们需要加入锁, 为了提高并发效率, 我们加入读写锁.
func main() {
	//初使化操作.
	u := UserAges{ages: map[string]int{}}
	wg := sync.WaitGroup{}
	//添加操作.
	for i := 0; i < 1000;i ++ {
		wg.Add(1)
		go func(i int) {
			u.Set("u"+strconv.Itoa(i), i)
			wg.Done()
		}(i)
	}
	//获取操作
	for i := 0; i < 1000;i ++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(u.Get("u"+strconv.Itoa(i)))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("finish")
}