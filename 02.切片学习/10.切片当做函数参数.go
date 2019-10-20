package main

import "fmt"

//切片是引用类型, 当做函数参数时,是一种指针类型传递
func Modify(s []int) { //实参相当于指针类型
	s[0] = 88//修改s指针指向的空间里的元素.原空间也是被修改
}
func main() {
	//声明s时,相当于s是一个指针类型,存储是切片的地址
	s := []int{1,2}
	fmt.Println(s)
	Modify(s)//传递时,相当是传递指针地址
	fmt.Println(s)
}
