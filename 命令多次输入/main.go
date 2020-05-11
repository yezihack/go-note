package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s2()
}

func s2() {
	type cmdS struct {
		No  int
		Msg string
	}
	cmdLine := []cmdS{
		{1, "生成表markdown文档"},
		{2, "生成表结构数据"},
		{3, "生成CURD增删改查"},
	}
	for _, row := range cmdLine {
		fmt.Printf("序列号:%d, %s\n", row.No, row.Msg)
	}
	fmt.Print("请输入以上命令序号:")
	r := bufio.NewReader(os.Stdin)
	c, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(c))
}

func s1() {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		fmt.Printf("Input Char Is : %v", string(input))
	}
}
