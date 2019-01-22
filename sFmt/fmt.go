package main

import (
	"fmt"
	"github.com/yezihack/studyGo/sFmt/s1"
	"runtime"
)

/*

   %+v 打印包括字段在内的实例的完整信息
   %#v 打印包括字段和限定类型名称在内的实例的完整信息
   %T 打印某个类型的完整说明
*/
func main() {
	s1.GetS1()
	fmt.Println(runtime.Caller(1))
	fmt.Println(runtime.Caller(2))
	fmt.Println(runtime.Caller(3))
	fmt.Println(runtime.Caller(4))
	aa := []int{1,2,3}
	fmt.Println(aa[2:])
}
