package simplemath

import (
	"github.com/yezihack/go-study/TheGoProgrammingLanguage/src"
	"testing"
)

func TestAdd(t *testing.T) {
	src.Asset(1, t, 3,  Add(1, 2))
	src.Asset(1, t, 14,  Add(12, 2))
	src.Asset(1, t, 15,  Add(13, 2))
}
