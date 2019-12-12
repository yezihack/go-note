package _2_锁

import (
	"runtime"
	"sync"
	"sync/atomic"
)

//自旋锁的实现
//利用原子操作.cas比较与交换
type spinLock int32

func (s *spinLock) Lock() {
	//未获取锁时, 循环检查等待, 直到获取锁,退出循环.
	for !atomic.CompareAndSwapInt32((*int32)(s), 0, 1) {
		runtime.Gosched()
	}
}
func (s *spinLock) Unlock() {
	atomic.StoreInt32((*int32)(s), 0)
}
//实例一个自旋锁.
func NewSpinLock() sync.Locker {
	var lock spinLock
	return &lock
}
