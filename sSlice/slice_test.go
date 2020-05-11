package sSlice

import (
	"fmt"
	"testing"
)

func TestStudySlice(t *testing.T) {
	StudySlice()
}

func BenchmarkForSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForSlice()
	}
}
func BenchmarkForArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForArray()
	}
}

func TestForSlice(t *testing.T) {
	SliceSort()
}

func TestLessSort(t *testing.T) {
	LessSort()
}

func TestMoon_Result(t *testing.T) {
	var ret string
	ret = moon.Result()
	ret = moon.Result()
	ret = moon.Result()
	fmt.Println(moon.len())
	ret = moon.Result()
	ret = moon.Result()
	ret = moon.Result()
	fmt.Println(moon.len())
	ret = moon.Result()
	ret = moon.Result()
	ret = moon.Result()
	fmt.Println(moon.len())
	ret = moon.Result()
	ret = moon.Result()
	fmt.Println(moon.len())
	fmt.Println(ret)
}
