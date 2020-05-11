package main

import "fmt"

func prime(n int) {
	set := make(map[int]bool, n)
	for i := 2; i < n; i ++ {
		set[i] = true
	}
	for i := 2; i <= n ; i ++ {
		if _, ok := set[i]; ok {
			fmt.Println(i)
			j := i + 1
			for j <= n {
				set[j] = false
				j += 1
			}
		}
	}
}
func main() {
	prime(10)
}
