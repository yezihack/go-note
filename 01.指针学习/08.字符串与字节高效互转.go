package main

import (
	"fmt"
	"unsafe"
)

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToByte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
func main() {
	//字节转字符串
	b := []byte{'a', 'b', 'c'}
	fmt.Println(ByteToString(b))
	//字符串转字节
	s := "abc"
	fmt.Println(StringToByte(s))
	//参考: https://studygolang.com/articles/2909
}
