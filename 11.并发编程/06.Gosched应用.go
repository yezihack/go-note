package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.Gosched 让出协程权给其它的协程.
// 让出当前协程资源 .
func main() {
	for i := 0; i < 5; i++ {
		//瞬间开辟5个子协程.
		go func(a int) {
			for j := 0; j < 3; j++ {
				if a == 2 {
					runtime.Gosched()
					//让出当前协程资源, 也就是说让这个资源往后放.
				}
				fmt.Printf("a:%d, i: %d, j:%d\n", a, i, j)
			}
		}(i)
	}
	time.Sleep(time.Second)
}
func main03() {
	for i := 0; i < 5; i++ {
		i := i
		//瞬间开辟5个子协程.
		go func(a int) {
			for j := 0; j < 3; j++ {
				if a == 2 {
					runtime.Gosched()
					//让出当前协程资源, 也就是说让这个资源往后放.
				}
				fmt.Printf("a:%d, i: %d, j:%d\n", a, i, j)
			}
		}(1)
	}
	time.Sleep(time.Second)
}

//改良版本.使用值传递给匿名函数打印.
func main02() {
	for i := 0; i < 2; i++ {
		//初使化goroutine需要消耗时间的.而goroutine里面的代码是初例化goroutine之后才执行的.
		go func(a int) {
			for j := 0; j < 10; j++ {
				fmt.Printf("i:%d, j:%d\n", a, j)
			}
		}(i)
	}
	time.Sleep(time.Second)
}
func main01() {
	for i := 0; i < 2; i++ {
		//初使化goroutine需要消耗时间的.而goroutine里面的代码是初例化goroutine之后才执行的.
		go func() {
			//for j := 0; j < 10; j++ {
			//	fmt.Printf("i:%d, j:%d\n", i, j)
			//}
		}()
	}
	time.Sleep(time.Second)
	/*
		会出现全是i=2,解决方法就是使用值传递.
		什么是值传递, 就是拷贝操作, 二个变量的地址是不一样的.
		i:2, j:4
		i:2, j:5
		i:2, j:6
		i:2, j:7
		i:2, j:8
		i:2, j:9
	*/
}
