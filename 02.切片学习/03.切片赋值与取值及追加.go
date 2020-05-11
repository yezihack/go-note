package main

import "fmt"

//切片赋值,取值, 追加

func main() {
	//初例一个0长度的切片,容量默认
	s := make([]int, 5)
	//赋值
	s[1] = 10
	fmt.Println(s)
	//取值
	val := s[1]
	fmt.Println(val)
	//追加使用append
	s = append(s, 100, 30)
	fmt.Println(s)
}
