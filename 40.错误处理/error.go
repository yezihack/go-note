package main

import "runtime"

// 自定义错误，打印堆栈信息

type stack []uintptr

type GoError struct {
	s string
	*stack
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[:n]
	return &st
}
func (e *GoError) Error() string {
	return e.s
}

func New(txt string) error {
	return &GoError{
		s:     txt,
		stack: callers(),
	}
}