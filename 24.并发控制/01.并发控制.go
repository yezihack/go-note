package _4_并发控制

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)
//并发控制

func Con() {
	//利用Context.WithTimeout设置超时
	cxt, cancel := context.WithTimeout(context.Background(), time.Second * 1)
	defer cancel()
	ch := make(chan int)
	go func() {
		time.Sleep(time.Millisecond * 1100)
		ch <- 132
	}()
	select {
	case val := <-ch:
		fmt.Println(val)
	case <-cxt.Done():
		fmt.Println("超时啦")
	}
}

//利用time.After
func ConAfter() {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Millisecond * 50)
		ch <- rand.Intn(100)
	}()
	select {
	case re := <-ch:
		fmt.Println("rand:", re)
	case <-time.After(time.Millisecond * 49):
		fmt.Println("超时")
	}
}