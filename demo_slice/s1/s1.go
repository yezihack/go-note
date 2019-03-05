package main

import (
	"fmt"
	"strings"
)

func main() {

}

const (
	MinCount = 100
	CentreCount = 1000
	MaxCount = 5000
	SCount = 5000
)

type Car struct {
	Name string
	Color string
	From string
	Age int
	Money float64
}


//错误写法,
func s0() {
	data := []int{1,2,3,4,5,6}
	for _, val := range data {
		fmt.Println(&val)
	}
}
// 输出结果:输出最后同一个地址
// 在执行for循环的时候，golang会首先创建一块内存，用于存放item。
// 之后依次将list中的元素拷贝到这块内存，在for之后若没有继续引用便进行释放，
// 所以在此过程中，修改item或将item放入其他的map中，只会放入最后一个元素。
/**
0xc420014060
0xc420014060
0xc420014060
0xc420014060
0xc420014060
0xc420014060
 */

 func tmpS0() {
 	data := []int{1,2,3,4,5}
 	for _, val := range data {
 		_ = &val
	}
 }
 func keyS1() {
	 data := []int{1,2,3,4,5}
	 for key := range data {
		 _ = &data[key]
	 }
 }

 //解决以上问题, 在for循环里每次创建一个val存储空间,这样就不会取同一块内容地址
func s1(data []int) {
	list := make([]*int, 0)
	for _, val := range data {
		val := val
		list = append(list, &val)
	}
}
//同样解决以上问题, 性能更优. 使用索引获取
func s2(data []int)  {
	list := make([]*int, 0)
	for key := range data {
		list = append(list, &data[key])
	}
}

//解决以上问题, 在for循环里每次创建一个val存储空间,这样就不会取同一块内容地址
func s3(data []Car) {
	for _, val := range data {
		val := val
		_ = &val
	}
}
//同样解决以上问题, 性能更优. 使用索引获取
func s4(data []Car)  {
	for key := range data {
		_ = &data[key]
	}
}


//非指针返回
func s5(data []Car) (result []Car) {
	for _, item := range data {
		result = append(result, item)
	}
	return
}
//指针返回
func s6(data []Car) (result []*Car) {
	for key := range data {
		result = append(result, &data[key])
	}
	return
}
//获取结构体数据
func GetData(count int) []Car {
	data := make([]Car, 0)
	for j := 0; j < count; j ++ {
		data = append(data, Car{
			Name:strings.Repeat("Sgfoot-Car", SCount),
			Color:strings.Repeat("white", SCount),
			Age: 10,
			Money: 100000.88,
		})
	}
	return data
}