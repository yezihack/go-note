package main

import "fmt"

func del(arr []int, idx int) []int {
	//先取出0:idx之间的数据, 然后再idx+1:最后的数据.再拼接.
	return append(arr[:idx], arr[idx+1:]...)
}
func main() {
	arr := []int{1,2,3,4,5}
	fmt.Println("del before:", arr)
	arr = del(arr, 1)
	fmt.Println("del after:", arr)

	fmt.Println("del before:", arr)
	arr = del(arr, 1)
	fmt.Println("del after:", arr)

	fmt.Println("del before:", arr)
	arr = del(arr, 1)
	fmt.Println("del after:", arr)

	fmt.Println("del before:", arr)
	arr = del(arr, 1)
	fmt.Println("del after:", arr)
}
