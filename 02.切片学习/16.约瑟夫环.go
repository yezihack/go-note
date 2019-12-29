package main

import "fmt"

func main() {
	fmt.Println(JosephRing(8, 3))
	fmt.Println(JosephRing(8, 4))
	fmt.Println(JosephRing(8, 5))
	fmt.Println(JosephRing(50, 3))
}
func JosephRing(n, k int) int {
	if n <= 0 || k <= 0 {
		return -1
	}
	ring := make([]int, n)
	for i := 0; i < n; i ++ {
		ring[i] = i + 1
	}
	record := 0 //record out count
	index := 0 //move array for index
	number := 0 //number off: 1,2,3,k
	for record < n - 1 {
		if ring[index] != 0 {
			number ++
		}
		if number == k {
			ring[index] = 0
			number = 0
			record ++
		}
		index ++
		if index == n {// if index to n else index = 0
			index = 0
		}
	}
	for i := 0; i < n; i ++ {
		if ring[i] > 0 {
			return ring[i]
		}
	}
	return -1
}
