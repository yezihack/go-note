package main

import "fmt"
//类型查询
func dump(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("%d", v.(int))
	case float64:
		fmt.Printf("%f", v.(float64))
	default:
		fmt.Println(v)
	}
}
func main() {
	dump(1)
}
