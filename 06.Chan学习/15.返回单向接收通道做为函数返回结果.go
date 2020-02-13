package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetNumber() <-chan int {
	c := make(chan int)
	go func() {
		c <- rand.Intn(100)
	}()
	return c
}
func Add(a, b int) int {
	return a + b
}
func main() {
	rand.Seed(time.Now().UnixNano())
	a, b := <-GetNumber(), <-GetNumber()
	fmt.Println(Add(a, b))
}
