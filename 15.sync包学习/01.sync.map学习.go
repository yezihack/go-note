/*
1. sync.map是并发安全的,开箱即用
适合于读多写少的场景, 为什么呢? 因为

2. 原理
引入2个map, 一个为读read map负责读, 不用加锁, 一个为写dirty map负责写,需要加锁.
当read map里没有读到值时,会加锁再读一次,如果还是没有读取,则dirty map上升为读map,记录miss++值
2个map底层数据是指针,指向同一份值
当miss大于等于dirty map时, dirty map会上升为read map, 这是会对dirty map进行地址拷贝


3.参考
https://juejin.im/post/5d74d562f265da03ab4273e1
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	sm := sync.Map{}
	sm.Store("aa", 11)
	v, ok := sm.Load("aa") //查找键是否存在
	if ok {
		fmt.Println("v", v)
	}
	//存储值,如果存在, OK=TRUE, 返回V1.返回二个元素.
	v1, ok := sm.LoadOrStore("aa", 22)
	if ok {
		fmt.Println("v1", v1)
	}
	//如果不存在, OK=FALSE
	v2, ok := sm.LoadOrStore("aa1", 22)
	fmt.Println("v2", v2, "ok", ok)
	sm.Range(func(key, value interface{}) bool {
		fmt.Println("range", "key", key, "value", value)
		return true
	})
}
