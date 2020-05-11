package main

import "fmt"

func main() {
	//声明一个指针,未分配空间,默认为nil
	var p *int
	//为指针分配一个有地址的空间
	p = new(int)
	//对指针指向的地址值修改
	*p = 300
	//打印指针指向的地址的值
	fmt.Println("*p=",*p)
	//打印指针的空间里的值,也就是地址
	fmt.Println("&p=", &p)
	//打印指向地址的值的空间地址
	fmt.Println("*&p=", *&p)
}
