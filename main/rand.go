package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
//map是并发不安全的.需要加锁操作.
const N = 10
func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	wg.Add(N)
	lock := sync.Mutex{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < N; i ++ {
		go func() {
			defer wg.Done()
			lock.Lock() //操作map需要加锁.否则会出并发不安全错误 .
			defer lock.Unlock()
			m[rand.Int()] = rand.Int()
		}()
	}
	wg.Wait()
	for key, val := range m {
		fmt.Printf("key:%d, value:%d\n", key, val)
	}
}
