package main

import (
	"testing"
	"strings"
)

//func BenchmarkS1(t *testing.B) {
//	for i := 0;i < t.N; i++ {
//		s1()
//	}
//}
//
//
//func BenchmarkS2(t *testing.B) {
//	for i := 0;i < t.N; i++ {
//		s2()
//	}
//}

const (
	COUNT = 50000
	JCount = 1000
)

func BenchmarkS3(t *testing.B) {
	t.ResetTimer()
	data := make([]Car, 0)
	for j := 0; j < JCount; j ++ {
		data = append(data, Car{
			Name:strings.Repeat("Benchi", COUNT),
			Color:strings.Repeat("white", COUNT),
			Age: 10,
			Money: 100000.88,
		})
	}
	t.StartTimer()
	for i := 0;i < t.N; i++ {
		s3(data)
	}
}
func BenchmarkS4(t *testing.B) {
	t.ResetTimer()
	data := make([]Car, 0)
	for j := 0; j < JCount; j ++ {
		data = append(data, Car{
			Name:strings.Repeat("Benchi", COUNT),
			Color:strings.Repeat("white", COUNT),
			Age: 10,
			Money: 100000.88,
		})
	}
	t.StartTimer()
	for i := 0;i < t.N; i++ {
		s4(data)
	}
}