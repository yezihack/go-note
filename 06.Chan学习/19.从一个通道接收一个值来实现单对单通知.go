package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		fmt.Println("A")
		t := time.NewTicker(time.Millisecond * 10)
		i := 0
		for {
			i ++
			fmt.Print(".")
			if i > 100 {
				break
			}
			<-t.C
		}
		<-done
	}()
	done <- struct{}{}
	fmt.Println("over!")
}
