package main

import "fmt"

func (c console) A() {
	fmt.Println("I'm console-A")
}

func (c console) B() {
	fmt.Println("I'm console-B")
}

func (c console) C() {
	fmt.Println("I'm console-C")
}

func (c child) AA() {
	fmt.Println("I'm Child AA")
}

func (c child) BB() {
	fmt.Println("I'm Child BB")
}