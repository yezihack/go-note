package main

import "fmt"

func InputNumber(s []int) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("请输入第%d个数:", i+1)
		_, err := fmt.Scanf("%d\n", &s[i])
		if err != nil {
			panic(err)
		}
	}
}
func Sum(s []int) int {
	sum := 0
	for _, num := range s {
		sum += num
	}
	return sum
}

func main() {
	//接收外部的一个元素,即大小
	//跟据大小输入多少个元素
	//然后相加,得到一个结果
	var Count, SumNumber int
	fmt.Println("请输入需要计算的个数:")
	_, err := fmt.Scanf("%d\n", &Count)
	if err != nil {
		panic(err)
	}
	s := make([]int, Count)
	InputNumber(s)
	SumNumber = Sum(s)
	fmt.Println("总共", SumNumber)

}
