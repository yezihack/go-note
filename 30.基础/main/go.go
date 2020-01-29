package main

import (
	"math"
	"sync"
)

func GSum(id int) {
	var x int64
	for i := 0; i < math.MaxInt32; i ++ {
		x += int64(i)
	}
	println(id, x)
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	for i := 0; i < 2; i ++ {
		go func(id int) {
			defer wg.Done()
			GSum(id)
		}(i)
	}
	wg.Wait()
}
