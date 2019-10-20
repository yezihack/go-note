package main

import "fmt"

func main() {
	var i, j int
	i, j = 3, 8
	fmt.Printf("交换前的值:i=%d, j=%d\n", i, j)
	swap(&i, &j)//进行地址传递.将i,j的地址传过去
	fmt.Printf("交换后的值:i=%d, j=%d\n", i, j)
}
//接收int类型的指针,也就是只接收地址传递.
func swap(i, j *int) {
	//将指针指向的地址的值进行交换
	*i, *j = *j, *i
}