package main

import "fmt"

//切片copy操作
//copy(目标切片s1, 需要拷贝的切片s2) ,也就是说s2向s1拷贝
func main() {
	s1 := []int{88, 99}
	s2 := []int{1,2,3,4,5}
	//第1种是长度小的切片向长度大的切片copy, s1 copy到s2切片里去
	copy(s2, s1)//将小的填到大的前面.进行覆盖
	fmt.Printf("s1:%v, s2:%v\n", s1, s2)//s1:[88 99], s2:[88 99 3 4 5]

	s1 = []int{55,66}
	s2 = []int{1,2,3,4,5}

	//第2种是长度大的切片向长度小的切片copy, s2 copy到s1切片里去
	copy(s1, s2)//最短的切片将被最长的切片覆盖, 长度为短的.
	fmt.Printf("s1:%v, s2:%v\n", s1, s2)//s1:[1 2], s2:[1 2 3 4 5]

	//修改s1, 会不会影响s2呢?
	s1[0] = 99 //不会影响,是一种拷贝,不是引用
	fmt.Printf("s1:%v, s2:%v\n", s1, s2)//s1:[99 2], s2:[1 2 3 4 5]
}
