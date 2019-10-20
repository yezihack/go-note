package main

import "fmt"

func main() {
	var p *int
	*p = 100
	fmt.Println(*p)
	//以上是声明了一个指针变量p,然后给指针变量指向的地址变量赋值赋值
	//而p指针没有地址,也就是p指针是nil
	//一个没有地址的指针,然后再给这个指针赋值的话就会panic
	//panic: runtime error: invalid memory address or nil pointer dereference

}
