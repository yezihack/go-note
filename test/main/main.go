package main

import (
	"fmt"
	"sync"
)

type Feed struct {
	wg *sync.WaitGroup
}

func (f *Feed) Print() {
	defer fmt.Println("feed+")
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("go ", err)
			}
		}()
		panic("go in panic")
	}()

}
func print_hello() {
	defer func() {
		fmt.Println("a")
	}()
	go func() {recover()}()
	//defer recover()
	defer func() {fmt.Println("b")}()
	panic("c")
}
func main() {
	print_hello()
}
