package test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSlice(t *testing.T) {
	a := make([]string, 0, 2)
	fmt.Printf("%p\n", a)
	App(a)
	fmt.Printf("%p\n", a)
	fmt.Println("a", a)
}
func App(a []string) {
	b := a
	fmt.Printf("a:%p, b:%p\n", a, b)
	for i := 0; i < 10; i++ {
		b = append(b, strconv.Itoa(i))
		fmt.Printf("a:%p, len:%d, cap:%d\n", a, len(a), cap(a))
		fmt.Printf("b:%p, len:%d, cap:%d\n", b, len(b), cap(b))
	}
	fmt.Println("b", b)
}
