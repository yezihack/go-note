package _0_基础

import (
	"fmt"
	"testing"
)

type Factory struct {
	Car
	Title string
}

type Car struct {
	Name string
}
func NewCar() *Car {
	return &Car{}
}
func (*Car) Show() {
	fmt.Println("car")
}
func NewFactory() *Factory {
	return &Factory{{"mini Car"}, "This is Factory"}
}
func TestCar(t *testing.T) {
	c := Car{}
	c.Show() //显式
	s := c.Show
	s()//隐式
}
