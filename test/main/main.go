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

func main() {
	var p *int
	p = new(int)
	*p = 300
	fmt.Println(*p)
	fmt.Println(&p)
	fmt.Println(*&p)
	fmt.Println(*&*&p)
}
