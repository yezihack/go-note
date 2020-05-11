package _0_基础

import (
	"fmt"
	"testing"
)
//不⽀支持 嵌套 (nested)、重载 (overload) 和 默认参数 (default parameter)
/*
•⽆无需声明原型。
•⽀支持不定⻓长变参。
•⽀支持多返回值。
•⽀支持命名返回参数。
•⽀支持匿名函数和闭包。
*/
//定义函数
type FormatFunc func(s string) int

func test(fn FormatFunc, s string) int {
	return fn(s)
}

func TestTt(t *testing.T) {
	le := test(func(s string) int {
		return len(s)
	}, "golang-php")
	fmt.Println(le)
}

//变参与多返回值
func funcA(x ...int) (int, string) {
	total := 0
	for _, a := range x  {
		total += a
	}
	return total, "success"
}
func TestFuncA(t *testing.T) {
	a, b := funcA(10, 2, 3, 4, 5)
	fmt.Println(a, b)
}
func Add(x, y int) (z int) {
	//var z = x + y //z 不能与返回的z同一个级别或作用域下定义. z redeclared in this block
	//return z
	{
		var z = x + y
		return z //不在同一个作用域下, 必须显式返回
	}
}
func TestAdd(t *testing.T) {
	fmt.Println(Add(10, 20))
}
//命名返回参数允许defer延时调用通过闭包读取和修改.
func Sum(x, y int) (z int) {
	defer func() {
		z += 100
	}()
	z = x + y
	fmt.Println("z", z)
	return z
}
func TestSum(t *testing.T) {
	fmt.Println(Sum(1, 2))
}
func Sum2(x, y int) (z int) {
	defer func() {
		fmt.Println("z", z)
	}()
	z = x + y
	return z + 200 // 执⾏行顺序: (z = z + 200) -> (call defer) -> (ret)
}
func TestSum2(t *testing.T) {
	fmt.Println(Sum2(1, 2))
}
//匿名函数可赋值给变量，做为结构字段，或者在 channel ⾥里传送。
func TestYM(t *testing.T) {
	fns := [](func(x int) int) {
		func(x int) int {
			return x + 1
		},
		func(x int) int {
			return x + 2
		},
	}
	fmt.Println(fns[1](10))
	//结构体里存放函数
	d := struct {
		fn func() string
	}{
		fn: func() string {
			return "golang"
		},
	}
	fmt.Println("d.fn():", d.fn())
	//chan里传输佚名函数
	fc := make(chan func() string, 2)
	fc <- func() string {
		return "php"
	}
	//fc返回一个函数, 必须<-fc括起来然后()
	fmt.Println("chan-func:", (<-fc)())
}
//闭包复制的是原对象指针，这就很容易解释延迟引⽤用现象。匿名
func BiBao() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}
func TestBiBao(t *testing.T) {
	f := BiBao()
	f()
}

func Div(x int) {
	defer fmt.Println("a")
	defer fmt.Println("b")
	//fmt.Println(10 /x) //b, a
	defer fmt.Println(10/ x)//b, a
	defer func() { //c, b, a
		fmt.Println(10 /x)
	}()
	defer fmt.Println("c")
}

//多个defer, 采用FILO先进后出
func TestDefer(t *testing.T) {
	Div(0)

}
func DeferTest() {
	x, y := 10, 20
	defer func(a int) {
		fmt.Println("defer:", a, y)// y 闭包引⽤用
	}(x)// x 被复制
	x += 10
	y += 20
	fmt.Println("x=", x, "y=", y)
}
func TestDeferTest(t *testing.T) {
	DeferTest()
}