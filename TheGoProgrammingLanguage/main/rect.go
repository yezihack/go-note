package main

import "fmt"

type Rect struct {
	x, y float64
	width, height float64
}
func (r *Rect) Area() float64 {
	return r.width * r.height
}
type Base struct {
	Name string
}
func (b *Base) Foo() {
	fmt.Println("base-foo", b.Name)
}
func (b *Base) Bar() {
	fmt.Println("base-bar",)
}
type Foo struct {
	//组合两个结构,具有相同的方法.不能直接调用,会出现模棱两可的错误.
	*Base
	Name string
}
func (foo *Foo) Bar() {
	foo.Foo() //正确的调用
	fmt.Println("foo-bar", foo.Name)
}

func main() {
	f := new(Foo)
	f.Name = "foo"
	f.Base = new(Base)
	f.Base.Name = "xxx"
	f.Bar()

}
