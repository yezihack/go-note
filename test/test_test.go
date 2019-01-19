package test

import (
	"fmt"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	Add(1000)
}
func BenchmarkFb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fb(20)
	}
}
func TestFb2(t *testing.T) {
	rs := Fb2(365)
	fmt.Println(rs)
}
func BenchmarkFb2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fb2(1000)
	}
}
