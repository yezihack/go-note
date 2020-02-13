package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getNum(c chan<- int) {
	c <- rand.Intn(100)
}
func AddNum(a, b int) int {
	return a + b
}
func main() {
	rand.Seed(time.Now().UnixNano())
	//申请了二个通道, 可以换成一个通道
	//c0, c1 := make(chan int), make(chan int)
	//go getNum(c0)
	//go getNum(c1)
	//fmt.Println(AddNum(<-c0, <-c1))

	//换成一个通道
	c := make(chan int)
	go getNum(c)
	go getNum(c)
	fmt.Println(AddNum(<-c, <-c))
}
