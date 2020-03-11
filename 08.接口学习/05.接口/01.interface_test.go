package _5_接口

import (
	"fmt"
	"testing"
)

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

func TestInterface(t *testing.T) {
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr
	i = &s2Val
	fmt.Printf("%v,%p\n", s1Val, s1Ptr)

	_ = s2Val
	_ = i
}
