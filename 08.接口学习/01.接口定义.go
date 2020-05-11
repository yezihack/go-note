package main

import (
	"fmt"
	"sync"
)

//go接口是一种非侵入式的,只需要实现接口里的方法就属于实现接口啦

type Cacher interface {
	Push(s interface{}) //添加方法
	Pop() interface{}   //弹出方法
}
type Cache struct {
	lock *sync.Mutex
	Data []interface{}
}

//初使cache
func NewCache() *Cache {
	return &Cache{
		lock: new(sync.Mutex),
		Data: make([]interface{}, 0),
	}
}

//添加方法
func (c *Cache) Push(s interface{}) {
	c.lock.Lock()
	c.Data = append(c.Data, s)
	c.lock.Unlock()
}
func (c *Cache) Pop() interface{} {
	if len(c.Data) == 0 {
		return nil
	}
	s := c.Data[0]
	c.Data = c.Data[1:]
	return s
}
func (c *Cache) Print() {
	for _, val := range c.Data {
		fmt.Println(val)
	}
}

func main() {
	var c Cacher
	c = NewCache()
	c.Push("aa")
	c.Push(100)
	c.Push(1.25)

	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())

	ch := NewCache()
	ch.Push(11)
	ch.Push("bb")
	ch.Print()

	//指明接口类型,只有实现接口里的所有的方法才能赋值成功.
	var cc Cacher = NewCache()
	cc.Push("pp")
	fmt.Println(cc.Pop())
}
