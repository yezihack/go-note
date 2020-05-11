package main

import "fmt"

//切片截取另一种 s[0:3:8]
func main() {
	s := []int{1,2,3,4,5}
	s1 := s[0:3:3] //0表示从0位置, 3表示到3的位置, 5表示s1的容量大小为5-0=5
	fmt.Printf("%v, %v, len:%d, cap:%d\n",s, s1, len(s1), cap(s1))//len:3, cap:5
	s2 := s[1:3:5] //// 切片截取 [low:high:max]  容量：max-low （max>=high,max<=len(原切片)）
	fmt.Printf("%v, %v, len:%d, cap:%d\n",s, s2, len(s2), cap(s2))//len:2, cap:4
}
