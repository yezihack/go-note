package main

import "fmt"

type Orange struct {
	Quantity int
}
func (o *Orange) Increase(n int) {
	o.Quantity += n
}
func (o *Orange) Decrease(n int)  {
	o.Quantity -= n
}
func (o *Orange) String() string {
	return fmt.Sprintf("%v", o.Quantity)
}
func main() {
	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange)
}
