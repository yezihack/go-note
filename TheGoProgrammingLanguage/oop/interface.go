package main

import "fmt"
//接口组合与查询
type Integer int
func (a Integer) Less(b Integer) bool {
	return a < b
}
//func (a *Integer) Less(b Integer) bool {
//	return (*a).Less(b)
//}
func (a *Integer) Add(b Integer) {
	*a += b
}

type Lesser interface {
	Less(b Integer) bool
	Add(b Integer)
}
type Lesser2 interface {
	Less(b Integer) bool
}

func main() {
	var a Integer = 10
	//var b Lesser = a
	var a2 Lesser = &a
	//fmt.Println(b.Less(12))
	fmt.Println(a2.Less(12))
	a2.Add(10)
	fmt.Println(a)

	var b Integer = 20
	var b1 Lesser2 = b
	var b2 Lesser2 = &b
	fmt.Println(b1.Less(1), b2.Less(2))
}
