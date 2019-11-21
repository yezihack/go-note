package main

import (
	"fmt"
	"sync"
	"time"
)

type SingleAtom struct {
}

var (
	sa   *SingleAtom
	once sync.Once
)

func NewSingleAtom() *SingleAtom {
	if sa == nil {
		once.Do(func() {
			sa = new(SingleAtom)
		})
	}
	return sa
}

func main() {
	go fmt.Printf("%p\n", NewSingleAtom())
	go fmt.Printf("%p\n", NewSingleAtom())
	go fmt.Printf("%p\n", NewSingleAtom())
	go fmt.Printf("%p\n", NewSingleAtom())
	go fmt.Printf("%p\n", NewSingleAtom())
	go fmt.Printf("%p\n", NewSingleAtom())
	go fmt.Printf("%p\n", NewSingleAtom())
	time.Sleep(time.Millisecond * 100)
}
