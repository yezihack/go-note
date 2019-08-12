package iota_demo

import (
	"fmt"
	"testing"
)

func TestShow(t *testing.T) {
	show()
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a[1:2:3])
}
