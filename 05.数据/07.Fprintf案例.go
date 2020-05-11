package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "wang\n")
	fmt.Fprintf(os.Stdout, "name:%s\n", "baili")

	var b []byte
	b = make([]byte, 3)
	b[0] = 'b'
	b[1] = 'b'
	b[2] = 'b'
	fmt.Println(b, b[:])
	n, err := os.Stdin.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("count:%d, string:%s\n", n, string(b[:]))
}
