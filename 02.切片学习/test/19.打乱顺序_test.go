package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// ShuffleSlice 打乱数组切片的顺序
func ShuffleSlice(slice []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
func Shuffle(s []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for n := len(s); n > 0; n-- {
		randIndex := r.Intn(n)
		fmt.Printf("n:%d, n-1:%d, rand:%d\n", n, n-1, randIndex)
		s[n-1], s[randIndex] = s[randIndex], s[n-1]
	}
}

func TestShift(t *testing.T) {
	a := []string{"a", "b", "c"}
	fmt.Println(a)
	//ShuffleSlice(a)
	//fmt.Println(a)
	//rand.Shuffle(len(a), func(i, j int) {
	//	a[i], a[j] = a[j], a[i]
	//})
	Shuffle(a)
	fmt.Println(a)
}
