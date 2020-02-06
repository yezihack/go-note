package _0_基础

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMemory(t *testing.T) {
	fmt.Println(unsafe.Sizeof(struct {
		i8 int8
		i16 int16
		i32 int32
	}{}))
	fmt.Println(unsafe.Sizeof(struct {
		i8 int8
		i32 int16
		i16 int32
	}{}))
}
