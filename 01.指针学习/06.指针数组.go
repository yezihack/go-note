package main

import "fmt"
//指针数组, 重点在数组,也就是说数组里每一个元素是指针
func main() {
	//定义一个指针数组, 就是说一个数组里的每一个元素都是存放指针地址的
	var p [3]*int
	var a = 10
	var b = 20
	var c = 30
	p[0] = &a //把a变量的地址给p数组里的第1位
	p[1] = &b
	p[2] = &c

	//打印第1个数组元素的值
	fmt.Println(*p[0])

	//修改指针数组
	*p[0] = 200
	fmt.Println(*p[0])
	Print2(p)
}

func Print2(nums [3]*int) {
	for key, val := range nums {
		//打印指针数组里的指向的空间里的值
		fmt.Println("第", key +1 , *val)
	}
}