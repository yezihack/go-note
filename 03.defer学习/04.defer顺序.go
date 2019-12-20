package main

import "fmt"

//go 的 defer 语句是用来延迟执行函数的，而且延迟发生在调用函数 return 之后

func main()  {
	fmt.Println(get())
	fmt.Println(get1())
	fmt.Println(*get2())
	a()
	fmt.Println("c", c())
}
func get() int {
	var i int
	defer func() {
		i ++
		fmt.Println("get")
	}()
	return i
}
func get1() (i int) {
	defer func() {
		i ++
		fmt.Println("get1")
	}()
	return i
}
func get2() *int {
	var i int
	defer func() {
		i ++
		fmt.Println("get2")
	}()
	return &i
}
func a() {
	i := 0
	defer fmt.Println(i)
	i ++
	return
}
func c() (i int) {
	defer func() {
		i ++
	}()
	return 1
}