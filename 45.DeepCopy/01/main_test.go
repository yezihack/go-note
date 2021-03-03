package main

import (
	"fmt"
	"testing"

	"github.com/mohae/deepcopy"
)

type Car struct {
	Name string
}

func TestDeepCopy(t *testing.T) {
	c1 := Car{
		Name: "Mini Car",
	}
	c2 := Car{}
	err := DeepCopy(&c1, &c2)
	if err != nil {
		t.Error(err)
	}
	c2.Name = "Bus Car"
	fmt.Printf("c1:%p, c1:%v\nc2:%p, c2:%v\n", &c1, c1, &c2, c2)
}
func TestDeepCopyReflect(t *testing.T) {
	c1 := Car{
		Name: "Mini Car",
	}
	c2 := deepcopy.Copy(&c1)
	switch c2.(type) {
	case *Car:
		cc2 := c2.(*Car)
		cc2.Name = "Bus Car"
	}

	fmt.Printf("c1:%p, c1:%v\nc2:%p, c2:%v\n", &c1, c1, &c2, c2)
}

func BenchmarkDeepCopy(b *testing.B) {
	c1 := Car{
		Name: "Mini Car",
	}
	c2 := Car{}
	for i := 0; i < b.N; i++ {
		DeepCopy(&c1, &c2)
	}
}

func BenchmarkDeepCopyReflect(b *testing.B) {
	c1 := Car{
		Name: "Mini Car",
	}
	for i := 0; i < b.N; i++ {
		deepcopy.Copy(&c1)
	}
}
func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
