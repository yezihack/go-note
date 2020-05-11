package main

import "fmt"

func main() {
	var i int
	i = 10
	str := "golang"
	m := make(map[string]bool)
	m["king"] = true
	//1.简单打印 %v, %s, %d
	fmt.Printf("%v\n", i) //打印变量的值, 可以是slice, array, int, float, string
	fmt.Printf("%s\n", str)//打印字符串
	fmt.Printf("%d\n", i)// 打印整型

	//2.打印%p, %T, %q, %x
	fmt.Printf("%p\n", m)//打印指针地址
	fmt.Printf("%T\n", i)//打印变量的类型
	fmt.Printf("%q\n", i)
	fmt.Printf("%x\n", i)//打印十六进制
	//3.详细打印, %+v, %#v
	fmt.Printf("%+v\n", m)
	fmt.Printf("%#v\n", m)
	//4.处理浮点数.
	f := 10.341
	fmt.Printf("%.2f", f)
}
