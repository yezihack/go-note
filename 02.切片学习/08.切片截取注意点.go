package main

import "fmt"

//切片是引用类型
func main() {
	s := []int{1,2,3, 4,5}
	s1 := s[1: 3]

	fmt.Println("修改前: s", s)
	fmt.Println("修改前: s1", s1)
	s1[0] = 100//修改s1的值, 会影响s的值
	fmt.Printf("修改后s:%v, p:%p\n", s, s)
	fmt.Printf("修改后s1:%v,p:%p\n", s1, &s1)

	//如果做到不影响呢?使用copy
	//s3 := make([]int,0)
	//copy(s3, s[1:3])
	//s3[1] = 200
	//fmt.Println("s", s)
	//fmt.Println("使用copy s3", s3)

}
