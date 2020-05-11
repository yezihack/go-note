package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

}
func h1() {
	http.HandleFunc("/ok", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("ok")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func getId(w http.ResponseWriter, r *http.Request) {

}
func h2() {
	s := &http.Server{
		Addr:":8080",
	}
	log.Fatal(s.ListenAndServe())
}