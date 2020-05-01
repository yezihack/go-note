package main

type Console interface {
	A()
	B()
	C()
}
type Child interface {
	AA()
	BB()
}

type console struct {
}
type child struct {
}

type Center interface {
	Console
	Child
}
type center struct {
	console
	child
}

func NewConsole() Center {
	return &center{}
}
func main() {
	c := NewConsole()
	c.A()
	c.B()
	c.C()
	c.AA()
	c.BB()
}