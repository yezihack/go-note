package main

import "fmt"

func main() {
	//声明一个普通的变量
	var i int = 10
	//声明一个指针变量,一个存放int类型的指针空间
	var p *int
	//把I变量的地址赋值给p指针
	p = &i
	//*p 是表示指针变量指向的变量的值.
	*p = 100
	//获取指针变量指向地址的值
	fmt.Println(*p)
}
