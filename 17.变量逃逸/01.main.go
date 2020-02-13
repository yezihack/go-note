package main

var x = 10
func foo(m0 int) (*int, *int) {
	var m1 int = 11
	var m2 int = 12
	var m3 int = 13
	var m4  = new(int)
	var m5  = new(int)
	*m4 = 5
	*m5 = 10
	println(&m0, &m1, &m2, &m3, &m4, &m5)
	println(&m0, &m1, &m2, &m3, &m4, &m5)
	println(&m0, &m1, &m2, &m3, &m4, &m5)
	println(&m0, &m1, &m2, &m3, &m4, &m5)
	println(&m0, &m1, &m2, &m3, &m4, &m5)
	//m1, m2, m4 分配在栈上.
	return &m0, m5 //m3, m5 从栈中逃逸出来, 到堆上.
}

func main() {
	//go tool compile -m 01.main.go
	m4, m5 := foo(100)
	println(*m4, *m5)
	println(x)
}

