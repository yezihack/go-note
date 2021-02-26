package _6_unsafe

import (
	"fmt"
	"testing"
)

func TestDouble(t *testing.T) {
	//a := 10
	//Double(&a)
	//fmt.Println(a)
	//p := &a
	//Double(p)
	//fmt.Println(a, p == nil)
}
func TestInt2float(t *testing.T) {
	a := 10
	f := Int2float(a)
	fmt.Printf("float:%f\n", f)
}
