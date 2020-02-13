package main

import "fmt"

func main() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default: //如果c缓冲已满,则执行default
		}
	}
	tryReceive := func() string {
		select {
		case v := <-c: return v
		default: return "nil" //如果c缓冲为空则执行default
		}
	}
	trySend("a") //send ok
	trySend("b") //send ok
	trySend("c") //send err , 但不会阻塞, 因为trysend 使用Select default处理了.
	//接收操作
	fmt.Println(tryReceive()) // a
	fmt.Println(tryReceive()) // b
	fmt.Println(tryReceive()) // nil
 }
