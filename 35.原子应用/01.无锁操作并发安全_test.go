package _5_原子应用

import "sync/atomic"

//原子的应用.
//实现一个并发安全不使用锁,只使用原子操作.
type ConcurrentArray interface {
	Set(index uint32, elem int) (err error)
	Get(index uint32) (elem int, err error)
	Len() uint32
}
type concurrentArray struct {
	length uint32
	val atomic.Value
}
func NewConcurrentArray(length uint32) ConcurrentArray {
	array := concurrentArray{}
	array.length = length
	array.val.Store(make([]int, array.length))
	return &array
}
//指针方法.
func (array *concurrentArray) Set(index uint32, elem int) (err error) {
	//todo
	return nil
}
func (array *concurrentArray) Get(index uint32) (elem int, err error) {
	//todo
	return 0, nil
}
func (array *concurrentArray) Len() uint32 {
	return array.length
}
