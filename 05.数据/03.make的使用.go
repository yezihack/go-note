/*
make只用于切片, 信道, map映射
使用slice, chan, map必须初使化,并非零值,返回一个指针 *T
以上三种都是引用类型.
*/
package main

import "fmt"

func main() {
	//复杂的声明, 先分配一个零值空间指针
	var ss *[]int = new([]int)
	//然后对值进行分配指向新的空间
	*ss = make([]int, 8)
	*ss = append(*ss, 2)
	fmt.Println(*ss)
	//简单的声明, 不返回指针, 返回一个引用类型对象.需要指针需要使用new
	s := make([]int, 0)
	s = append(s, 44)
	fmt.Println(s, &s)
}
