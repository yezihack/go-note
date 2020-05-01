package main

import "fmt"
//切片可以理解成: 动态数组
//切片通过对数组进行封装，为数据序列提供了更通用、强大而方便的接口
func main() {
	//1.声明一个切片
	var s []int
	s = append(s, 10)
	fmt.Printf("len: %d, list: %v\n", len(s), s)

	//2.简单推法声明
	s1 := []int{1}
	s1 = append(s1, 20)
	fmt.Printf("len: %d, list: %v\n", len(s1), s1)

	//3.make([]类型, 长度, 容量) 长度不能小于0
	s3 := make([]int, 3, 8) //定义了一个长度为3(默认为0值), 空间为8的切片
	fmt.Printf("len: %d, list: %v\n", len(s3), s3)

	//4.推导并实例值
	s4 := []int{1,2,3}
	fmt.Printf("len: %d, list: %v\n", len(s4), s4)

	//5.添加多个值
	s5 := []int{1,2 }
	s5 = append(s5, 3, 4,5)
	fmt.Printf("len: %d, list: %v\n", len(s5), s5)

	//注意1. 如果切片大小为0,不能使用s[0] = 10操作的
}
type abc struct {
	
}