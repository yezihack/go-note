package _0_基础

import (
	"errors"
	"fmt"
	"testing"
)
//没有结构化异常，使⽤用 panic 抛出错误，recover 捕获错误
//由于 panic、recover 参数类型为 interface{}，因此可抛出任何类型对象。
func PanicA() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()
	panic("panic error")
}

func TestPanic(t *testing.T) {
	PanicA()
}
//延迟调⽤用中引发的错误，可被后续延迟调⽤用捕获，但仅最后⼀一个错误可被捕获。
func TestPanicB(t *testing.T) {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	fmt.Println("bb")
	panic("test panic")
}
//捕获函数 recover 只有在延迟调⽤用内直接调⽤用才会终⽌止错误，否则总是返回 nil。任何未
//捕获的错误都会沿调⽤用堆栈向外传递
func TestPanicC(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer recover() // ⽆无效！
	defer fmt.Println(recover()) // ⽆无效！
	defer func() {
		func() {
			println("defer inner")
			recover() // ⽆无效！
		}()
	}()
	panic("test panic")
}
func CatchPanic() {
	recover()
}
//使⽤用延迟匿名函数或下⾯面这样都是有效的。
func TestC1(t *testing.T) {
	//defer CatchPanic()
	defer func() {
		recover()
	}()
	panic("C panic")
}
//revover 必须使用, 放在函数里头.才能恢复错误.

//如何区别使⽤用 panic 和 error 两种⽅方式？惯例是：导致关键流程出现不可修复性错误的
//使⽤用 panic，其他使⽤用 error。

var ErrDivZero = errors.New("divsion by zero")
func ErrDiv(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivZero
	}
	return x / y, nil
}
func TestErrDiv(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic:%#v, %s", err, err.(error).Error())
		}
	}()
	switch z, err := ErrDiv(10, 0); err {
	case nil:
		fmt.Println(z)
	case ErrDivZero:
		panic(err)
	}
}