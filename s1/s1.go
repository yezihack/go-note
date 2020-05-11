package main

import "fmt"

func main() {
	s0()
}

//错误写法,
func s0() {
	data := []int{1, 2, 3, 4, 5, 6}
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

//解决以上问题, 在for循环里每次创建一个val存储空间,这样就不会取同一块内容地址
func s1() {
	data := []int{1, 2, 3, 4, 5, 6}
	list := make([]*int, 0)
	for _, val := range data {
		val := val
		list = append(list, &val)
	}
}

//同样解决以上问题, 性能更优. 使用索引获取
func s2() {
	data := []int{1, 2, 3, 4, 5, 6}
	list := make([]*int, 0)
	for key := range data {
		list = append(list, &data[key])
	}
}

//基准测试结果. S2 > S1 相当2倍性能.
/**
pkg: demo_slice/s1
BenchmarkS1-4            5000000               266 ns/op             168 B/op         10 allocs/op
BenchmarkS2-4           10000000               256 ns/op             168 B/op          5 allocs/op
PASS
ok     demo_slice/s1     4.398s
*/

type Car struct {
	Name  string
	Color string
	From  string
	Age   int
	Money float64
}

//解决以上问题, 在for循环里每次创建一个val存储空间,这样就不会取同一块内容地址
func s3(data []Car) {
	for _, val := range data {
		val := val
		_ = &val
	}
}

//同样解决以上问题, 性能更优. 使用索引获取
func s4(data []Car) {
	for key := range data {
		_ = &data[key]
	}
}
