package main

import (
	"fmt"
	"sync"
)

func main() {
	demo1()
}
func demo1() {
	type value struct {
		mu    sync.Mutex
		value int
	}

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()         //1
		defer v1.mu.Unlock() //2

		//time.Sleep(2 * time.Second) //3
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
func demo() {
	var data int
	var mt sync.Mutex
	go func() {
		mt.Lock()
		defer mt.Unlock()
		data ++
	}()
	mt.Lock()
	defer mt.Unlock()
	if data == 0 {
		fmt.Println("value ", data)
	} else {
		fmt.Println("value", data)
	}
}