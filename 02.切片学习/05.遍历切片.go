package main

import (
	"fmt"
	"math/rand"
	"time"
)

//遍历切片
func main() {
	s := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <5; i ++  {
		s = append(s, rand.Intn(100))
	}
	//遍历
	for key, val := range s {
		fmt.Println("key", key, "value", val)
	}
}
