package main

import (
	"fmt"
)

func Print2(i int) {
	fmt.Println(i)
}

func main() {
	for i := 0; i < 1000000; i++ {
		go Print2(i)
	}
	fmt.Println("over")
}
