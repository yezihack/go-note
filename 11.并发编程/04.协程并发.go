package main

import "fmt"

func Print() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	//将print函数放在子协程里执行.
	//开辟新的协程执行.与下面的for是同步执行的.
	go Print()
	//主进程执行
	for i := 90; i < 100; i++ {
		fmt.Println(i)
	}
	fmt.Println("main over")
	//主程序挂掉, 子协程必挂
}
