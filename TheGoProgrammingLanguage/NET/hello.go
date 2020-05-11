package main

import (
	"io"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
	//buf := bytes.NewBufferString("hello world!!")
	//w.Write(buf.Bytes())
}
func main() {
	http.HandleFunc("/hello", helloWorld)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
