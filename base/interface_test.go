package _base

import (
	"fmt"
	"testing"
)

type F interface {
	f()
}
type S1 struct {
}

func (s S1) f() {
	//panic("implement me")
}

type S2 struct {
}

func (s *S2) f() {
	//panic("implement me")
}
func TestF(t *testing.T) {
	s1Val := S1{}
	s1Ptr := &s1Val
	s2Val := S2{}
	s2Ptr := &s2Val
	var i F
	i = s1Val
	i = s1Ptr
	//S2 does not implement F (f method has pointer receiver)
	//i = s2Val
	i = s2Ptr
	fmt.Print(i)
}
