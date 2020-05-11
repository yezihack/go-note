package main

type mspan struct {
	addr uintptr
}
func main() {
	var a *int
	*a = 10
	m := new(mspan)
	m.addr = a
}
