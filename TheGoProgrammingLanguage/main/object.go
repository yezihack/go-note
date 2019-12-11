package main

import (
	"fmt"
)

type Integer int
func (i Integer) Less(b Integer) bool {
	return i < b
}
func (i *Integer) Add(b Integer) {
	*i += b
}

func main() {
	var a Integer = 10
	fmt.Println(a.Less(11))
	a.Add(12)
	fmt.Println(a)

}
