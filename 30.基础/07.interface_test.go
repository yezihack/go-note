package _0_基础

import (
	"fmt"
	"testing"
)

type Man struct {
	Id int
	Name string
}
func (self *Man) String() string {
	return fmt.Sprintf("%d %s", self.Id, self.Name)
}

func TestInterface(t *testing.T) {
	var o interface{} = &Man{1, "Tom"}
	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println("i", i)
	}
	u := o.(*Man)
	fmt.Println(u.Id, u.Name)
}
func TestInterfaceSwitch(t *testing.T) {
	var o interface{} = &Man{2, "KOB"}
	switch o.(type) {
	case nil:
		fmt.Println("o is nil")
	case *Man:
		fmt.Println("o is Man")
	case fmt.Stringer:
		fmt.Println("o is fmt.Stringer")
	case func() string:
		fmt.Println("o is func()")
	default:
		fmt.Println("unknown")
	}
}

//某些时候，让函数直接 "实现" 接⼝口能省不少事。
type Tester interface {
	Do()
}
type FuncDo func()
func (self FuncDo) Do() {
	self()
}
func TestFuncDo(t *testing.T) {
	var do Tester = FuncDo(func() {
		fmt.Println("my is Do")
	})
	do.Do()
}