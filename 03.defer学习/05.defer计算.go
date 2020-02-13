package main

import "fmt"

func main() {
	i := 10
	defer func() {
		i += 10
		fmt.Println("D", i)
	}()
	fmt.Println("A", i)
	defer func() {
		i += 10
		fmt.Println("C", i)
	}()
	fmt.Println("B", i)
	/*
A 10
B 10
C 20
D 30
	*/
}
