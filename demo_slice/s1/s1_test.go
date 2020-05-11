package main

import (
	"fmt"
	"runtime"
	"testing"
)

/**
基准测试 循环时获取有效地址的两种操作
一种是使用临时变量
一种是使用索用获取
效率深入分析

使用方法: ./compile.sh 运行得知结果
*/

/*************************小数据, 简单切片操作, 大约100个数据*********************************/
func init() {
	fmt.Println("马立全开")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Benchmark小数据100_简单切片操作0(t *testing.B) {
	t.ResetTimer()
	data := make([]int, 0)
	for j := 0; j < MinCount; j++ {
		data = append(data, j)
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		s1(data)
	}
}

func Benchmark小数据100_简单切片操作1(t *testing.B) {
	t.ResetTimer()
	data := make([]int, 0)
	for j := 0; j < MinCount; j++ {
		data = append(data, j)
	}
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		s2(data)
	}
}

/*
小数据的指针切片赋值(一个使用临时变量,一个使用索引) s1 和 s2 简单的基准测试
Benchmark小数据100_简单切片操作0-4                                500000              2311 ns/op            2840 B/op        108 allocs/op
Benchmark小数据100_简单切片操作1-4                               2000000               839 ns/op            2040 B/op          8 allocs/op
性能相差4倍左右
推荐指针操作时使用:索引获取地址
*/

/*******************************小数据, 结构操作,大约100个******************************/

func Benchmark小数据100_结构操作0(t *testing.B) {
	t.ResetTimer()
	data := GetData(MinCount)
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		s3(data)
	}
}
func Benchmark小数据100_结构操作1(t *testing.B) {
	t.ResetTimer()
	data := GetData(MinCount)
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		s4(data)
	}
}

/**
小数据结构体操作

Benchmark小数据100_结构操作0-4                           2000000               699 ns/op               8 B/op          0 allocs/op
Benchmark小数据100_结构操作1-4                          20000000                68.5 ns/op             0 B/op          0 allocs/op
性能相差10倍左右
推荐指针操作时使用:索引获取地址
*/
/*************************中数据, 结构体操作, 大约1000个数据*********************************/

func Benchmark中数据1000_结构操作0(t *testing.B) {
	t.ResetTimer()
	data := GetData(CentreCount)
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		s3(data)
	}
}
func Benchmark中数据1000_结构操作1(t *testing.B) {
	t.ResetTimer()
	data := GetData(CentreCount)
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		s4(data)
	}
}

/**
中数据结构体操作

Benchmark中数据1000_结构操作0-4                           200000              7069 ns/op             846 B/op          0 allocs/op
Benchmark中数据1000_结构操作1-4                          2000000               592 ns/op              84 B/op          0 allocs/op
性能相差10倍左右
推荐指针操作时使用:索引获取地址
*/

/*************************大数据, 结构体操作, 大约5000个数据*********************************/

func Benchmark大数据5000_结构操作0(t *testing.B) {
	t.ResetTimer()
	data := GetData(MaxCount)
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		s3(data)
	}
}
func Benchmark大数据5000_结构操作1(t *testing.B) {
	t.ResetTimer()
	data := GetData(MaxCount)
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		s4(data)
	}
}

/**
大数据结构体操作

Benchmark大数据5000_结构操作0-4                            30000             42540 ns/op           28253 B/op          0 allocs/op
Benchmark大数据5000_结构操作1-4                           300000              3508 ns/op            2825 B/op          0 allocs/op
性能相差10倍左右
推荐指针操作时使用:索引获取地址
*/

/*************************指针与非指针数据返回, 结构体操作, 大约5000个数据*********************************/

func Benchmark非指针结构5000_返回操作0(t *testing.B) {
	t.ResetTimer()
	data := GetData(MaxCount)
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		_ = s5(data)
	}
}
func Benchmark有指针结构5000_返回操作1(t *testing.B) {
	t.ResetTimer()
	data := GetData(MaxCount)
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		_ = s6(data)
	}
}

/**
指针与非指针数据返回操作

Benchmark非指针结构5000_返回操作0-4                                 3000            408176 ns/op         1797987 B/op         24 allocs/op
Benchmark有指针结构5000_返回操作1-4                                30000             93699 ns/op          185173 B/op         17 allocs/op
性能相差10倍左右
推荐指针操作时使用:索引获取地址
*/
