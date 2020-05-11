package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func printPrime(prefix string) {
	defer wg.Done()
next:

	for outer := 2; outer < 50; outer ++ {
		for inner := 2; inner < outer; inner++ {
			if outer % inner == 0 {
				fmt.Printf("out:%d, inner:%d\n", outer, inner)
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed:", prefix)
}
func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go printPrime("A")
	go printPrime("B")
	wg.Wait()
}
