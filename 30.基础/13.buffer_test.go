package _0_基础

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	var b bytes.Buffer
	fmt.Fprint(&b, "abc")
	fmt.Println(b.String())
}
