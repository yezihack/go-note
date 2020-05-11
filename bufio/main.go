package main

import (
	"bufio"
	"os"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)
//参考: https://studygolang.com/articles/4367

func main() {
	inputString()
}
// ReadString 功能同 ReadBytes，只不过返回的是一个字符串
func inputString() {
	s, err := bufio.NewReader(os.Stdout).ReadString(' ')
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	spew.Dump(s)
}
//获取设置长度的字符,过度则丢掉
func inputByte() {
	b := make([]byte, 10)//10个字节
	n , err := bufio.NewReader(os.Stdout).Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len: %d, s: %s\n", n, string(b))
	spew.Dump(b)
}
//ReadLine 尝试返回一个单行数据，不包括行尾标记（\n 或 \r\n）
func inputByteLine() {
	line , prefix, err := bufio.NewReader(os.Stdout).ReadLine()
	if err != nil {
		panic(err)
	}
	fmt.Printf("len: %d, prefix: %v, s: %s\n", len(line), prefix, string(line))
	spew.Dump(line)
}
