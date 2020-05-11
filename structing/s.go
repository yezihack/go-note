package structing

import "fmt"

type Person struct {
	Name string
}

func New(name string) *Person {
	return &Person{
		Name: name,
	}
}

var std = New("默认")

func ShowName() {
	fmt.Println("Person Name:", std.Name)
}
