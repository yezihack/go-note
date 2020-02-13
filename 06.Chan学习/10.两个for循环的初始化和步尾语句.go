package main

import (
	"fmt"
	"time"
)

func main() {
	fab := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			var x, y  uint64 = 0, 1
			for ; y < (1 << 63); c <- y {
				x, y = y, x + y
			}
			close(c)
		}()
		return c
	}
	c := fab()
	//for x, ok := <-c; ok; x, ok = <-c {
	//	time.Sleep(time.Millisecond *100)
	//	fmt.Println(x)
	//}
	for x := range c {
		time.Sleep(time.Millisecond *100)
		fmt.Println(x)
	}

}
