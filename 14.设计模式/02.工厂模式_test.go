package Patterm

import (
	"fmt"
	"testing"
)

func TestNewFly(t *testing.T) {
	f := NewFly(WARPLANE)
	fmt.Println("什么飞机:", f.Effect())
	fmt.Println("轮子数量:",f.Wheel())

	f1 := NewFly(COPTER)
	fmt.Println("什么飞机:", f1.Effect())
	fmt.Println("轮子数量:",f1.Wheel())
}
