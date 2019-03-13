package sSlice

import "testing"

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
