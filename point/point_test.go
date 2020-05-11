package point

import (
	"testing"
)

func TestStruct(t *testing.T) {
	var a, b struct{}
	println(&a)
	println(&b)
	println(&a == &b, &a == nil)
	if a != b {
		t.Error("not same")
	}
	if &a != &b {
		t.Error("point address not same")
	}
}
func TestSWitch(t *testing.T) {
	switch x := 5; {
	case x > 5:
		println(5)
	case x > 0 && x <= 5:
		println(0, 5)
	default:
		println("2")
	}
}
func TestForWithString(t *testing.T) {
	a := "golang"
	for x, v := range a {
		println(x, "-", v, "->", string(v))
	}
	x := 0
	for x < 10 {
		x++
		println(x)
	}
}
