package _6_unsafe

import (
	"fmt"
	"unsafe"
)

func ChangeByte() {
	x := 0x12345678
	fmt.Printf("%T\n", x)
	p := unsafe.Pointer(&x)
	n := (*[4]byte)(p)
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X, %T \n ", n[i], n[i])
	}
}
