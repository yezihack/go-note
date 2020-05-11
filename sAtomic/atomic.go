package sAtomic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func StudyAtomic() {
	var sum uint32
	var sum2 uint32
	var wg sync.WaitGroup
	for i := 0; i < 999; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum2 += 1
			atomic.AddUint32(&sum, 1)
		}()
	}
	wg.Wait()
	fmt.Println("线程安全,计算的sum变量::", sum)
	fmt.Println("非线程安全,计算的sum变量:", sum2)
}
