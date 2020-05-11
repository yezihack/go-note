package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	//recover只能在当前的goroutine发挥作用.
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("go", err)
			}
		}()
		panic("A")
	}()
	time.Sleep(time.Second)
	fmt.Println("ccc")
	//go A
	//ccc
}
