package main

import (
	"fmt"
	"github.com/ivpusic/grpool"
	"time"
)

func main() {
	pool := grpool.NewPool(10, 5)
	defer pool.Release()
	for i := 0; i < 100; i++ {
		count := i
		pool.JobQueue <- func() {
			fmt.Printf("I am worker! Number %d\n", count)
		}
	}
	time.Sleep(time.Second * 1)
}
