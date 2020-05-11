package main

import "fmt"

type Footer interface {
	string()
}

type Personer interface {
	Footer
	test()
}

type person struct {
}

func (p *person) string() {
	fmt.Println("string")
}
func (p *person) test() {
	fmt.Println("test")
}

func foot(f Footer) {
	f.string()
}
func main() {
	var p person
	var pp Personer = &p
	foot(pp) //隐式转换为子集接口

	pp.string()
	pp.test()

	var f Footer = pp //超集转换为子集
	f.string()

	//var p2 Personer = f 超集转换子集可以, 返过来不可以

}
