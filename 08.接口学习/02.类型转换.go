package main

import "fmt"

//类型断言. T.(type)
func main() {
	//1. 复杂写法.借用switch
	var i interface{}
	i = 10
	fmt.Println(CheckType(i))
	s := []int{2, 3, 3}
	fmt.Println(CheckType(s))
	//2. 简单写法
	if val, ok := i.(int); ok {
		fmt.Println("整型", val)
	}
}
func CheckType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case []int:
		return "[]int"
	default:
		return "other"
	}
}
