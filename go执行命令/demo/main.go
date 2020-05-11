package main

import (
	"fmt"
	"os/exec"
)

func test() {
	cmd := exec.Command("sh", "-c", "nslookup www.baidu.com")
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

func main() {
	go1()
}
func go1() {
	cmd := exec.Command("goimports1", "-l", "-w", "/Users/wangzl/go-work/src/github.com/yezihack/gm2m/output/structure")
	bytes, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
