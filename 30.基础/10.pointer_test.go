package _0_基础

import (
	"fmt"
	"testing"
	"unsafe"
)
//除正常指针外，指针还有 unsafe.Pointer 和 uintptr 两种形态。其中 uintptr 被 GC 当
//做普通整数对象，它不能阻止所 "引用" 对象被回收
func TestPointer(t *testing.T) {
	var u uintptr
	var x unsafe.Pointer
	fmt.Println(u, x)
}