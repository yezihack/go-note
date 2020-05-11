package main

import "fmt"

func foo() *int {
	a := 2
	return &a
}
func coo() int {
	b := 3
	return b
}

func main() {
	x := foo()
	fmt.Println(*x)
	y := coo()
	fmt.Println(y)
	/*
	使用命令: go build -gcflags "-m -l" main.go 进行分析变量逃逸情况
	go tool compile -S 01.escapes.go
	*/
}

