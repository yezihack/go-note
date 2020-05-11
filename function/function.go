package function

import "fmt"

type N int

func (n N) toString() string {
	return fmt.Sprintf("%#x", n)
}
func (n *N) Pointer() {
	fmt.Printf("p:%p", n)
}

func show() {
	var i N = 10

	fmt.Println(i.toString())
	i.Pointer()
}

func call(m func(n int), n int) {
	m(n)
}
func display(n int) {
	n++
	fmt.Println("n:", n)
}
func Hello() {
	call(display, 2)
	call(display, 10)
}
