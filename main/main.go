package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello world")
	})
	fmt.Println("start :9090")
	http.ListenAndServe(":9090", nil)
}
