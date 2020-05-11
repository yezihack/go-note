package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func setup() {
	fmt.Println("hello world")
}
func doPrint() {
	once.Do(setup)
}
func main() {
	go doPrint()
	go doPrint()
	go doPrint()
	go doPrint()
	go doPrint()
	time.Sleep(time.Millisecond * 100)
}
