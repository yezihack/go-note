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
	type s struct {
		a int
		b int
	}
	var b = s{1, 2,
		}
	
}
