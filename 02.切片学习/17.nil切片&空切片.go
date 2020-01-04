package main

import "fmt"

func main() {
	//nil 切片和空切片很相似，长度和容量都是0，官方建议尽量使用 nil 切片。
	var s1 []int
	var s2 = []int{}
	var s3 = make([]int, 0)
	var s4 = *new([]int)
	fmt.Println("s1",  s1 == nil) //true
	fmt.Println("s2",  s2 == nil) //false
	fmt.Println("s3",  s3 == nil) //false
	fmt.Println("s4",  s4 == nil) //true
	s1 = make([]int, 0)
	s1 = append(s1, 1)
	fmt.Println("s1", s1)

	s2 = append(s2, 2)
	fmt.Println("s2", s2)

	s3 = append(s3, 3)
	fmt.Println("s3", s3)

	s4 = make([]int, 0)
	s4 = append(s4, 4)
	fmt.Println("s4", s4)
}