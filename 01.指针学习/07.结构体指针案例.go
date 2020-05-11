package main

import "fmt"

//结构体指针案例
type Student struct {
	ID int
	Name string
}
//传递一个指针变量做为实参.传递的是一个指针, 修改结构体会影响传递的指针变量
func ModifyByPoint(student *Student) {
	student.ID = 99
	student.Name = "王也道长"
}
//传递一个结构体变量,是拷贝形式传递, 会产生副本,修改不会影响其它值
func Modify(student Student) {
	student.ID = 80
	student.Name = "冯宝宝"
}
func main() {
	s := &Student{100, "百里"}
	Modify(*s)
	fmt.Println(*s)//{100 百里}
	ModifyByPoint(s)
	fmt.Println(*s)//{99 王也道长}
}
