package main

import (
	"os"
	"fmt"
	"net"
	"time"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	cmdFile := os.Args[0]
	fmt.Println("文件名:", cmdFile)
	firstName := os.Args[1]
	fmt.Println("第1个参数:", firstName)

	conn, err := net.DialTimeout("tcp", firstName, time.Duration(time.Second * 3))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	spew.Dump(conn)
}
