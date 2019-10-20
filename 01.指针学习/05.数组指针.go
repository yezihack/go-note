package main

import "fmt"
//数组指针,重点是指针, 这个数组是指针.
func main() {
	//定义一个数组
	var nums [3]int
	nums = [3]int{1,2,3}
	//定义一个存储数组变量的指针
	var p *[3]int //这是一个数组指针, 就是说他是一个数组类型的指针
	//将数组的地址赋值给指针p
	p = &nums
	//获取p的值,即地址
	fmt.Println(p)
	//获取p指向的空间,即数组的值
	fmt.Println(*p)
	//获取指向空间的数组第一个元素
	fmt.Println((*p)[0])
	//修改指向空间的数组第2个元素
	(*p)[1] = 200
	//获取所有的数组值
	fmt.Println(*p)
	Print(&*p)//这二个是相同的
	Print(&nums)//这二个是相同的
}

func Print(arr *[3]int) {
	fmt.Println("-------------")
	fmt.Println((*arr)[0])
	fmt.Println((*arr)[1])
	fmt.Println((*arr)[2])
	for _, val := range *arr {
		fmt.Println(val)
	}
}