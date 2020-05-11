package _9_对象池

import (
	"fmt"
	"sync"
	"testing"
)

//sync.pool
//保存和复用临时对象，减少内存分配，降低GC压力。
//GC会清除sync.pool缓存的对象
//对象的缓冲生命周期是GC下一次调用之前


func TestSyncPool(t *testing.T) {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	fmt.Println("p1", p.Get().(int))
	p.Put(10)
	fmt.Println("p2", p.Get().(int)) //获取10
	fmt.Println("p3", p.Get().(int)) //获取0,因为已经被gc回收啦
	fmt.Println("p4", p.Get().(int)) //获取0,因为已经被gc回收啦

}