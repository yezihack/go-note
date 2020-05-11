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
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("go", err)
			}
		}()
		panic("aaa")
	}()
	time.Sleep(time.Second)
}
