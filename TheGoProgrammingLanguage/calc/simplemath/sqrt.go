package simplemath

import "math"

func Sqrt(a int) int {
	return int(math.Sqrt(float64(a)))
}