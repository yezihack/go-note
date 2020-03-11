package _5_接口

import (
	"fmt"
	"testing"
)
type Animal interface {

}

type Car struct {
	name string
}

func (c Car) ShowName() {
	fmt.Println(c.name)
}
func (c *Car) SetName(name string) {
	c.name = name
}

func TestName(t *testing.T) {
	c := map[int]Car{1: {name:"bus"}}//值对象，只能调用值方法，不能调用指针方法
	//c[1].SetName("mini bus")//cannot call pointer method on c[1]
	c[1].ShowName()
	c1 := map[int]*Car{1:{name:"car"}}//指针对象是可以调用值方法或指针方法
	c1[1].SetName("mini car")
	c1[1].ShowName()
}
