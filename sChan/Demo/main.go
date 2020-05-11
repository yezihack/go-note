package main

import (
	"net/http"
	"fmt"
)

var c chan int

func main() {
	c = make(chan int, 3)
	go func() {
		for {
			select {
			case d := <-c:
				fmt.Println(d)
			}
		}
	}()
	http.HandleFunc("/in", In)
	http.ListenAndServe(":8080", nil)
}
func In(w http.ResponseWriter, r *http.Request) {
	c <- 10
}
