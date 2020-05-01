package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var tag int
	flag.IntVar(&tag, "tag", 1, "tag 1 or 2")
	flag.Parse()
	switch tag {
	case 1:
		Sel()
	case 2:
		select2()
	}
}
var g chan int
var quit chan bool
func select2() {
	g = make(chan int)
	quit = make(chan bool)
	go B()
	<-time.After(time.Second * 3)
	quit <- true // 没办法等待B的退出只能Sleep
	fmt.Println("Main quit")
}

func B() {
	for {
		select {
		//case i := <-g:
		//	fmt.Println(i + 1)
		case <-quit:
			fmt.Println("B quit")
			return
		}
	}
}

func Sel() {
	ch := make(chan int)
	go func() {
		t := time.NewTicker(time.Second * 3)
		select {
		case <- t.C:
			ch <- 20
		}
	}()
	select{
	case s := <-ch:
		fmt.Println("ok", s)
	}
}
