package main

import (
	"fmt"
)

func main() {
	mutex := make(chan struct{}, 1)
	couter := 0
	increase := func() {
		mutex <- struct{}{}
		couter ++
		<-mutex
	}
	i100 := func(done chan<- struct{}) {
		for i := 0; i < 100; i ++ {
			increase()
		}
		done <- struct{}{}
	}
	done := make(chan struct{}, 100)
	for i := 0; i < cap(done);i ++ {
		go i100(done)
	}
	for i := 0; i < cap(done);i ++ {
		<-done
	}
	fmt.Println(couter)
}
