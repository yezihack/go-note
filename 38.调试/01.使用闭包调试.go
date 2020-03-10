package main

import (
	"fmt"
	"runtime"
	"time"
)

func m1() {
	fmt.Println("m1")
}
func m2() {
	fmt.Println("m2")
}
func m3() {
	fmt.Println("m3")
}
func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
	}
	sAt := time.Now()
	where()
	m1()
	where()
	time.Sleep(time.Millisecond * 100)
	m2()
	where()
	m3()
	eAt := time.Now()
	result1 := eAt.Sub(sAt)
	result2 := time.Since(sAt)
	fmt.Printf("r:%s, r2:%s\n", result1, result2)
	fmt.Println(sAt.Before(eAt)) //sAt在eAt之前，返回true
}