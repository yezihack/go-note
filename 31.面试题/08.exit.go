package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("A")
	}()
	fmt.Println("ABC")
	//0代表成功, 非0代表错误.
	//defer不会执行的.
	os.Exit(0)
	defer func() {
		fmt.Println("B")
	}()

}
