package main

import "fmt"

type Person struct {
	ID int
	Name string
	c chan int
}
//对结构体使用new初使化零值空间,里面的chan并没有初使化.需要额外实例.
//&Person 与 New(Person)是等价的
func main() {
	p := new(Person)
	p.ID = 100
	p.c = make(chan int)
	fmt.Println(p.ID)
	go func() {
		p.c <- 10
	}()
	fmt.Println(<-p.c)
}
func AddPerson(Id int, name string) *Person {
	return &Person{
		Name:name,
		ID:Id,
	}
	//&Person 与 New(Person)是等价的
}
