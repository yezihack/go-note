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
func (*Car) Show() {
	fmt.Println("car")
}
func NewFactory() *Factory {
	return &Factory{
		 Car{"Mini Car"},
		"aa",
	}
}
func TestCar(t *testing.T) {
	c := Car{}
	c.Show() //显式
	s := c.Show
	s()//隐式
}
