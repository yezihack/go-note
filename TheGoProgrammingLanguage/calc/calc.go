package main

import (
	"fmt"
	"github.com/yezihack/go-study/TheGoProgrammingLanguage/calc/simplemath"
	"log"
	"os"
	"strconv"
)
var Usage = func() {
	fmt.Println("USAGE:calc comand [arguments] ...")
}

func main() {
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	args := os.Args[1:]
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	switch args[0] {
	case "add":
		if len(args) != 3 {
			log.Println("USAGE:calc add <integer1> <integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		v2, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			log.Println("USAGE:calc add <integer1> <integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("result", ret)
	case "sqrt":
		if len(args) != 2 {
			log.Println("USAGE:calc sqrt <integer>")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		if err1 != nil {
			log.Println("USAGE:calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v1)
		fmt.Println("result:", ret)
	default:
		Usage()
	}
}
