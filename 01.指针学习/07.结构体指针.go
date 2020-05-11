package main

import "fmt"

//结构体指针, 重点是指针. 也就是说指针指向一个结构体

type Person struct {
	Name string
	Sex int
}
func main() {
	//将结构体的地址赋值给p结构体指针.
	var p *Person = &Person{
		Name:"百里",
		Sex:1,
	}
	//简写
	//p := &Person{"百里", 1}
	//修改结构体指针里的成员
	(*p).Name = "三百里"
	//简写 p.Name = "三百里"
	fmt.Println("p.Name=", p.Name)

	person := Person{"王也", 1}
	fmt.Println(person.Name)
}
