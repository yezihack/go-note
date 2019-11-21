package main

import (
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func (f HandleFunc) ServeHttp(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
func main() {

}
