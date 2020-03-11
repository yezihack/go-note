package _5_接口

import (
	"fmt"
	"testing"
)

type Cuper interface {
	Display()
}
type Cup struct {
}
func (c *Cup) Display() {
	fmt.Println("cup")
}
func Show(s Cuper) {
	s.Display()
}

type CC struct {
}
func (c CC) Display() {
	fmt.Println("cc")
}
func TestCup(t *testing.T) {
	c := Cup{}
	//Show(c) //Cup没有值接收器。必须使用地址传递。
	Show(&c)
	cc := CC{}
	Show(cc)
	Show(&cc)
}