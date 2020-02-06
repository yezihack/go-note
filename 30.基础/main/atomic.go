package main

import (
	"fmt"
	"sync/atomic"
)

func FuncValue(v atomic.Value) {
	v.Store([]int{1,2,3})
}
func main() {
	var a atomic.Value
	a.Store([]int{8,9})
	FuncValue(a)
	fmt.Println(a.Load())
}