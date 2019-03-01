package demo2

import (
	"fmt"
	"github.com/yezihack/studyGo/理解defer执行顺序/common"
)

func init() {
	defer func() {
		fmt.Println( common.Now(), "我是demo2文件夹里的demo2_1.go文件里的init")
	}()
}