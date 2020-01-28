package _0_基础

import (
	"fmt"
	"testing"
)

func switchIf(x int) {
	switch {
	case x > 1:
		fmt.Println("x > 1")
		fallthrough
	case x > 5:
		fmt.Println("x > 5")
	case x < 0:
		fmt.Println("x < 0")
	default:
		fmt.Println("x", x)
	}
}
func switchInit(x []int) {
	switch x[0] = 10; {
	case x[0] > 10:
		fmt.Println("x>10")
	case x[0] < 0:
		fmt.Println("x < 0")
	default:
		fmt.Println("x", x[0])
	}
}

func TestSwitch(t *testing.T) {
	switchIf(10)
	a := []int{2}
	switchInit(a)
}
