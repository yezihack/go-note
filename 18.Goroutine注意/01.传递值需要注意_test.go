package _8_Goroutine注意

import (
	"fmt"
	"testing"
	"time"
)

func TestValue(t *testing.T) {
	//第一种正确写法
	fmt.Println("第一种正确写法")
	for i := 0; i < 10; i ++ {
		go func(x int) {
			fmt.Print(x, ",")
		}(i)
	}
	//正确的结果:9,4,3,0,2,7,8,5,6,1,
	//这是方法调用都是值传递.保证了变量的安全
	time.Sleep(time.Millisecond * 50)
	fmt.Println()
	//第二种错误写法
	fmt.Println("第二种错误写法")
	for j := 0; j < 10; j ++{
		go func() {
			fmt.Print(j, ",")
		}()
	}
	//结果:10,10,10,4,10,10,10,10,10,8, 出现了重复的数字.
	// 这是因为在当前main协程里与 for j 里的协程共享了j变量,导致有资源的竞争
	time.Sleep(time.Millisecond * 50)
}
