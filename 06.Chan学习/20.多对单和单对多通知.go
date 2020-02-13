package main

import (
	"log"
	"time"
)

type T = struct {}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready
	log.Println("Work#", id, "start Work")
	time.Sleep(time.Second * time.Duration(id+1))
	log.Println("work#", id, "end Work")
	done <- T{}//通知主协程（N-to-1）
}
func main() {
	log.SetFlags(0)
	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)
	time.Sleep(time.Millisecond * 500)
	// 单对多通知
	//ready<- T{}
	//ready<- T{}
	//ready<- T{}
	close(ready) //close(ready) // 群发通知
	// 等待被多对单通知
	<-done
	<-done
	<-done
}
