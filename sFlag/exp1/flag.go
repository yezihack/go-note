package main

import (
	"flag"
	"fmt"

	f "github.com/yezihack/studyGo/sFlag/exp1/flag"
)

func main() {
	var tables string
	flag.StringVar(&tables, "table", "", "表名称,多个表名使用逗号隔开")
	flag.Parse()
	if tables == "" {
		fmt.Println("请带个参数")
		return
	}
	fmt.Println(tables)
	fmt.Println(f.GetName())
}
