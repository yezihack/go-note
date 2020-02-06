package main

import (
	"fmt"
	"math"
)

func max(a, b int64) int64 {
	return int64(math.Max(float64(a), float64(b)))
}
func main() {
	fmt.Println(max(1, 2))
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MaxInt64-1, float64(math.MaxInt64-1))
	fmt.Println(math.MaxInt64-2)
	fmt.Println(max(math.MaxInt64-1, math.MaxInt64-2))
}
