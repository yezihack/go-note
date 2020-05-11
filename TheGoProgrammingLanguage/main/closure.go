package main

import "fmt"

func main() {
	var j = 5
	a := func() func() {
		var i = 3
		return func() {
			fmt.Printf("i:%d, j:%d\n", i, j)
		}
	}()
	a()
	j += 2
	a()
}
