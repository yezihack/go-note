package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var chain string
func main() {
	chain = "main"
	A()

	fmt.Println(chain)
}
func A() {
	mu.Lock()
	defer mu.Unlock() //必须调用B函数,而B函数又调用C函数,C函数里又加锁啦, 两次加锁.导致死锁.deadlock;
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}
