package demo1

import (
	"fmt"

	"github.com/yezihack/studyGo/理解go执行顺序/common"
)

func init() {
	defer func() {
		fmt.Println(common.Now(), "我是demo1文件夹里的demo1.go文件里的init")
	}()
}

func Show(i int) {
	defer func() {
		fmt.Printf("%v 我是demo1文件夹里的demo1.go文件里的show函数, number :%d\n", common.Now(), i)
	}()
}
