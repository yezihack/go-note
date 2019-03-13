package demo2

import (
	"fmt"

	"github.com/yezihack/studyGo/理解defer执行顺序/common"
)

func Show(i int) {
	defer func() {
		fmt.Printf("%v 我是demo2文件夹下的demo2.go文件里的show方法, number: %d\n", common.Now(), i)
	}()
}
