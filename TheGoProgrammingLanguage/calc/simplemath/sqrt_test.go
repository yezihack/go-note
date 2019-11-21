package simplemath

import (
	"github.com/yezihack/go-study/TheGoProgrammingLanguage/src"
	"testing"
)

func TestSqrt(t *testing.T) {
	src.Asset(1, t, 2, Sqrt(4))
	src.Asset(1, t, 3, Sqrt(9))
}
