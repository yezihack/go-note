package main

import "fmt"

func A() error {
	return New("This is A")
}
func B() error {
	return A()
}
func C() error {
	return B()
}
func main() {
	err := C()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
