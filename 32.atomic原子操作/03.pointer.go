package _2_atomic

import (
	"sync/atomic"
	"time"
	"unsafe"
)
//使用自旋获取对象. 保证只有一个对象.
type EtcdClient struct {
	username string
	password string
	id int
}
func NewEtcdClient() *EtcdClient {
	return &EtcdClient{
		username:"etcd",
		password:"root!root",
		id: 100,
	}
}
var etcdClientPointer unsafe.Pointer

func Client() *EtcdClient {
	for {
		p := (*EtcdClient)(atomic.LoadPointer(&etcdClientPointer)) //如果指针存在则直接返回
		if p != nil {
			return p
		}
		//不存在则创建
		c := NewEtcdClient()
		//使用比较并载入操作
		if !atomic.CompareAndSwapPointer(&etcdClientPointer, nil, unsafe.Pointer(&c)) {
			time.Sleep(time.Nanosecond * 10) //休10纳秒
			continue
		}
		return c
	}
}
