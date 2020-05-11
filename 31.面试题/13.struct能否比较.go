package main

import "fmt"

type a struct {
	name string
}
type b struct {
	name string
}
func main() {
	i := a{name:"aa"}
	j := a{name:"bb"}
	if i == j {
		fmt.Println("ok")
	} else {
		fmt.Println("err")
	}
	s1 := make([]int, 10)
	s2 := make([]int, 33)
	if &s1 == &s2 {
		fmt.Println("ok")
	} else {
		fmt.Println("err")
	}
}