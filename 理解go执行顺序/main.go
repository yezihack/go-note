package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/yezihack/studyGo/理解go执行顺序/common"
	"github.com/yezihack/studyGo/理解go执行顺序/demo1"
	"github.com/yezihack/studyGo/理解go执行顺序/demo2"
)

func init() {
	fmt.Println(common.Now(), "我是main.go里的init")
}

func main() {
	number := flag.Int("n", 0, "0为Normal, 1为GoChannel")
	flag.Parse()
	if *number == 0 {
		Normal()
	} else {
		GoChannel()
	}
}
func Normal() {
	fmt.Println("Normal执行中....")
	fmt.Println(common.Now(), "我是main函数 START")
	demo1.Show(1)
	demo2.Show(2)
	time.Sleep(time.Millisecond * 500)
	fmt.Println(common.Now(), "我是main函数 END")
}
func GoChannel() {
	fmt.Println("GoChannel执行中....")
	fmt.Println(common.Now(), "我是main函数 START")
	go demo1.Show(1)
	go demo2.Show(2)
	time.Sleep(time.Millisecond * 500)
	fmt.Println(common.Now(), "我是main函数 END")
}
