package _0_基础

import (
	"fmt"
	"testing"
)

//•数组是值类型，赋值和传参会复制整个数组，⽽而不是指针。
//•数组⻓长度必须是常量，且是类型的组成部分。[2]int 和 [3]int 是不同类型。
//•⽀支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
//•指针数组 [n]*T，数组指针 *[n]T

func TestArrayInit(t *testing.T) {
	a := [3]int{1,2,3}
	b := [...]int{33, 21, 2}
	c := [5]int{1:10, 2:20}
	d := [...]struct{
		name string
		age int
	}{
		{name:"baili", age:22},
		{name:"baili", age:23},
		{name:"baili", age:24},
	}
	fmt.Println(a, b, c, d)
}