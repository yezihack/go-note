package main

import "fmt"

type bigint int64
func main() {
	x := 1024
	var b = bigint(x)
	fmt.Println(b)
}


