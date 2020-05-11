package main

import "fmt"

type Parent struct {
}
func (p *Parent) MethodB() {
	fmt.Println("method-B")
}
func (p *Parent) MethodA() {
	fmt.Println("method A")
	p.MethodB()
}
type Child struct {
	Parent
}
func (c *Child) MethodB() {
	fmt.Println("child - method - B")
}
func (c *Child) MethodA() {
	fmt.Println("child method A")
	c.MethodB()
}
func main() {
	child := Child{}
	child.MethodA()
	child.Parent.MethodA()
}
