package main

import (
	"fmt"
	"math/rand"
	"runtime/debug"
	"time"
)

func fast(c chan<- int) {
	s := rand.Intn(5)
	fmt.Println("产生一个", s, "s")
	time.Sleep(time.Duration(s) * time.Second)
	c <- rand.Int()
}

func main() {
	defer debug.PrintStack()
	rand.Seed(time.Now().UnixNano())
	startAt := time.Now()
	c := make(chan int, 10)
	for i := 0; i < cap(c); i ++ {
		go fast(c) //开多个go
	}
	rnd := <-c // 只有第一个回应被使用了
	fmt.Println(time.Since(startAt))
	fmt.Println(rnd)
}
