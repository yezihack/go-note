package main

import "fmt"

func main() {
	var i = 10
	var p *int = &i
	*p += 10
	fmt.Println(i, &p, &i)

}
