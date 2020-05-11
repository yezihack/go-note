package main

import "fmt"

type Singleton struct {
}

var singleton *Singleton

//这种线程不安全, 并发会产生二个对象
func NewSingleton() *Singleton {
	if singleton == nil {
		singleton = new(Singleton)
	}
	return singleton
}
func main() {
	for i := 0; i < 100; i++ {
		go func() {
			addr := NewSingleton()
			fmt.Printf("%p\n", addr)
		}()
	}
}
