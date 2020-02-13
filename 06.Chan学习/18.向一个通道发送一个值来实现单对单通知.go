package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	values := make([]byte, 32 * 1024 * 1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	done := make(chan struct{}) //做信号通知用的
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{}
	}()
	<-done
	//fmt.Println(values)
	fmt.Println(values[0], values[len(values)-1])

}
