package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile|log.LstdFlags)
	testGo2Div()
	time.Sleep(time.Millisecond * 500)
}
func testGo2Div() {
	go func() {
		defer func() {
			//会将div里的panic扑捉到.程序不会奔溃.但是goroutine里的panic是抓不到的.
			//解决办法就是在goroutine里写上recover
			if r := recover(); r !=nil {
				log.Println("Runtime error caught", r)
			}
		}()
		div(10, 0)
	}()
}
func testGoDiv() {
	defer func() {
		//会将div里的panic扑捉到.程序不会奔溃.但是goroutine里的panic是抓不到的.
		//解决办法就是在goroutine里写上recover
		if r := recover(); r !=nil {
			log.Println("Runtime error caught", r)
		}
	}()
	go div(10, 0)
}
func testDiv() {
	defer func() {
		//会将div里的panic扑捉到.程序不会奔溃.
		if r := recover(); r !=nil {
			log.Println("Runtime error caught", r)
		}
	}()
	div(10, 0)
}

func div(a, b int) int {
	if b == 0 {
		panic("b is 0")
	}
	return a/b
}