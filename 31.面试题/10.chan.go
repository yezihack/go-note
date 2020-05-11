package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()
		for _, value := range set.s {
			ch <- value
		}
		close(ch)
	}()
	return ch
}
func SendRand(ch chan<- int) {
	ch <- rand.Intn(10)
}
func main() {
	ch := make(chan int, 1)
	SendRand(ch)
	fmt.Println(<-ch)
}
