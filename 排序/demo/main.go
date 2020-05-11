package main

import (
	"fmt"
	"sort"
)

type Car struct {
	Name string
}
type CarList []Car

func (c CarList) Len() int { //求长度
	return len(c)
}
func (c CarList) Less(i, j int) bool { //比大小
	return c[i].Name < c[j].Name //增序 [{apple} {board} {keep} {zero}]
	//return c[i].Name > c[j].Name //降序 [{zero} {keep} {board} {apple}]
}
func (c CarList) Swap(i, j int) { //交换位置
	c[i].Name, c[j].Name = c[j].Name, c[i].Name
}

func main() {
	method1()
	fmt.Println("--------------------------")
	method2()
}

//方法二
//调用 sort.Slice() 排序 ,简单,代码少
func method2() {
	fmt.Println("方法二: 调用 sort.Slice() 排序")
	c := make([]Car, 0)
	c = append(c, Car{
		Name: "QQ",
	})
	c = append(c, Car{
		Name: "Audi",
	})
	c = append(c, Car{
		Name: "VOLVO",
	})
	fmt.Println("排序前:", c)
	sort.Slice(c, func(i, j int) bool {
		return c[i].Name < c[j].Name
	})
	fmt.Println("排序后:", c)

}

//排序方法一
//使用额外的三个方法实现
func method1() {
	fmt.Println("方法一: 使用额外的三个方法实现")
	a := CarList{
		Car{
			Name: "zero",
		},
		Car{
			Name: "apple",
		},
		Car{
			Name: "board",
		},
		Car{
			Name: "keep",
		},
	}
	sort.Sort(a)
	fmt.Println(a)
}
