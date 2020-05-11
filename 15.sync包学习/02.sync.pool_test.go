package sync_study

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

//sync.pool
//保存和复用临时对象，减少内存分配，降低GC压力。
//底层原理: 私有对象, 共享池

//获取get & 写入put
//GET获取. 先从私有对象里获bc,如果没有则在共享链表里查找,如果还没有则调用new
//获取: 先会去私有对象里去获取,这里是线程安全的,
// 不存在则去共享池里,线程不安全,则会加锁获取.保证资源安全.
// 如果共享池里也不存在,则会调用new对象.
//写入: 先向私有对象存放, 不存在保存为私有对象.当前私有对象存在,则放入当前的processor子池的共享池中.
//共享池里是一个链表, 写入头部.

// 生命周期
//GC会清除sync.pool缓存的对象
//对象的缓冲生命周期是GC下一次调用之前


func TestSyncPool(t *testing.T) {

	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a New Object")
			return 100
		},
	}
	//直接获取, 私有对象里不存在, 共享里也不存在,则直接调用new
	fmt.Println(pool.Get())
	//向pool里存放一个值
	pool.Put(200)
	pool.Put(300)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	runtime.Gosched()
}
func TestSyncPoolGet(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} {
			return 2
		},
	}
	fmt.Println(p.Get())
}