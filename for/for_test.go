package _for

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	a := []int{
		1, 2, 3,
	}
	for k, v := range a {
		if k == 0 {
			a[0] *= 10
			a[1] *= 10
			a[2] *= 10
		}
		fmt.Printf("val:%d, data:%d\n", v, a[k])
	}
}

func TestFor2(t *testing.T) {
	a := []int{
		1, 2, 3,
	}
	for k, v := range a[:] {
		if k == 0 {
			a[0] *= 10
			a[1] *= 10
			a[2] *= 10
		}
		fmt.Printf("val:%d, data:%d\n", v, a[k])
	}
}
