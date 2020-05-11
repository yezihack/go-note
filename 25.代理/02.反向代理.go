package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/gommon/log"
)

type handle struct {
	host string
	port string
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + h.host + ":" + h.port)
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func main() {
	h := &handle{
		host: "localhost",
		port: "9091",
	}
	//proxy
	err := http.ListenAndServe(":89", h)
	if err != nil {
		log.Fatal(err)
	}
	//e := echo.New()
	//e.Server.Handler = h
	//e.Server.Addr = ":89"
	//err := e.Server.ListenAndServe()
	//if err != nil {
	//	log.Fatal(err)
	//}

}
