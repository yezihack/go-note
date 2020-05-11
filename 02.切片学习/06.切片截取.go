package main

import (
	"fmt"
)

//切片截取
func main() {
	s := []int{1,2,3,4,5,6,7,8,9}
	//s[i] i = 0 获取下标为0的值
	fmt.Println("单个元素:",s[0])//结果:1
	//s[0:j] 简写为 s[:j] 从0位置到j的位置(不包含j位置的值)
	fmt.Println("截取前部分:",s[:2])//结果: 1, 2
	//s[i:len(s)-1],简写:s[i:]从i的位置到最后一个元素
	fmt.Println("截取后部分",s[2:])//结果:3,4,5,6,7,8,9
	//s[i:j] 从i的位置开始(包含i位置的元素), 到j的位置结束(不包括j位置的元素)
	fmt.Println("截取指定中间部分:",s[1: 3])//结果:2, 3
	//s[:] 表示取所有的值
	fmt.Println("所有的值:", s[:])

	s1 := s[1:3] //从下标1到3的位置截取. 容量就是之前的9-1=8
	fmt.Printf("len:%d, cap:%d\n", len(s), cap(s))//len:9, cap:9
	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1))//len:2, cap:8
}
