package main

import (
	"fmt"
	"unsafe"
)

func main() {
	b := make([]byte, 10)
	b[0] = 'a'
	b[1] = 'b'
	s := *(*string)(unsafe.Pointer(&b))
	fmt.Println(s)
}
func StringToBytes(s string) byte {
	return *(*byte)(unsafe.Pointer(&s))
}
func ByteToString(b byte) string{
	return *(*string)(unsafe.Pointer(&b))
}