package test

import "testing"
//切片（slice）是对数组一个连续片段的引用，所以切片是一个引用类型
func ForArray() {
	array := [1024]int{}
	for i :=0;i < 1024; i ++ {
		array[i] = i
	}
}
func ForSlice() {
	s := make([]int, 0)
	for i :=0;i < 1024; i ++ {
		s = append(s, i)
	}
}

func BenchmarkArray(t *testing.B) {
	for i := 0; i <t.N; i ++{
		ForArray()
	}

}
func BenchmarkSlice(t *testing.B) {
	for i := 0; i <t.N; i ++{
		ForSlice()
	}
}
//以下结论, 在操作数组时比切片更快.数组每次执行堆上分配内存总量是0，分配次数也是0 。
//go test -bench . -benchmem -gcflags "-N -l"
//BenchmarkArray-8          500000              2422 ns/op               0 B/op          0 allocs/op
//BenchmarkSlice-8          300000              6119 ns/op           16376 B/op         11 allocs/op
//PASS

