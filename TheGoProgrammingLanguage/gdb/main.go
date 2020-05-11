package main

import "fmt"

func main() {
	 a, b := 10, 20
	 c := add(a, b)
	 fmt.Println(c)
}
func add(a, b int) int{
	return a + b
}
