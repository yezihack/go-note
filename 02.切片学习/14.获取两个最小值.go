package main

import (
	"fmt"
)

//一个循环里找出二个最小值.

func main() {
	a := []int{102, 21, 1, 3, 9, 99, 33, 4}
	fmt.Println(a)
	fmt.Println(min2(a))
	fmt.Println(min3(a))

}
func min2(a []int) (min1, min2 int) {
	min1, min2 = 1<<32, 1<<32
	for _, val := range a {
		if min1 > val {
			min2 = min1 //100
			min1 = val  //8
			//fmt.Printf("min1>val: min:%d, min2:%d, val:%d\n", min1, min2, val)
		} else if min2 > val { //8
			min2 = val
			//fmt.Printf("min2>val: min:%d, min2:%d, val:%d\n", min1, min2, val)
		}
	}
	return
}

func min3(a []int) (m1, m2, m3 int) {
	m1 = 1 << 32
	m2, m3 = m1, m1
	for _, v := range a {
		if m1 > v {
			m2 = m1
			m1 = v
		} else if m2 > v {
			m3 = m2
			m3 = v
			m2 = v
		} else if m3 > v {
			m3 = v
		}
	}
	return
}
