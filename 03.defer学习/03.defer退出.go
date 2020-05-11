package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	os.Exit(0)
}