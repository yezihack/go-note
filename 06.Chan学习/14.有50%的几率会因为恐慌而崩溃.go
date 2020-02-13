package main

import "fmt"

func main() {
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}://存在50%的机率,产生panic
	case <-c:
		fmt.Println("nil")
	}
}
