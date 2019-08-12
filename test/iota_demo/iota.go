package iota_demo

import "fmt"

const (
	a = iota
	b = iota
	c = 10
	d = iota + 1
	e
)

func show() {
	fmt.Println(e)
}
