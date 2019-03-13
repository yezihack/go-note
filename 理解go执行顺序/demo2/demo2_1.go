package demo2

import (
	"fmt"

	"github.com/yezihack/studyGo/理解go执行顺序/common"
)

func init() {
	fmt.Println(common.Now(), "我是demo2文件夹里的demo2_1.go文件里的init")
}
