package main

import (
	"fmt"
	"sync"
)

type Single struct {
}

var s *Single
var mu sync.Mutex

func NewSingle() *Single {
	if s == nil {
		mu.Lock()
		defer mu.Unlock()
		s = new(Single)
	}
	return s
}

func main() {
	go fmt.Printf("%p\n", NewSingle())
	go fmt.Printf("%p\n", NewSingle())
	go fmt.Printf("%p\n", NewSingle())
	go fmt.Printf("%p\n", NewSingle())
	go fmt.Printf("%p\n", NewSingle())
	go fmt.Printf("%p\n", NewSingle())
}
