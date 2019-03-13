package main

import "fmt"

type Person struct {
	Age int
}
var c = make(chan int, 1)

func New() Person {
	p := Person{}
	go func() {
		for  {
			select {
			case d := <-c:
				p.Age = d
			}
		}
	}()

	return p
}
//
//func (p *Person) SetAge(age int) {
//	p.Age = age
//}

func (p *Person) Print() {
	fmt.Printf("%d", p.Age)
}


func main() {
	c <- 900
	p := New()
	p.Print()
}

