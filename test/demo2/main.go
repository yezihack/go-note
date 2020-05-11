package main

import "fmt"

func mkmap() map[string]int {
	m := make(map[string]int)
	m["a"] = 1
	return m
}
func mkslice(a int) []int {
	s := make([]int, 0, 10)
	s = append(s, 1)
	return s
}

func main() {
	m := mkmap()
	fmt.Println(m["a"])

	s := mkslice()
	fmt.Println(s[0])
}
