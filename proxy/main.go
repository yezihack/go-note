package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handle struct {
	host string
	port string
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + h.host + ":" + h.port)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}
func StartServer() {
	h := &handle{
		host: "127.0.0.1",
		port: "8000",
	}
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		panic(err)
	}
}
func main() {
	StartServer()
}
