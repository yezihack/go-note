package main

import (
	"fmt"
	"os"
)

func main() {
	//获取环境变量
	path := os.Getenv("GOPATH")
	fmt.Println(path)
}
