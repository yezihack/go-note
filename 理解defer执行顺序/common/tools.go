package common

import (
	"time"
	"fmt"
)

func init() {
	defer func() {
		fmt.Println(Now(), "我是common文件夹里的tools.go文件中的init")
	}()
}

func Now() time.Time {
	chinaTime, _ := time.LoadLocation("Asia/Chongqing")
	time.Local = chinaTime
	return time.Now().Local()
}




























