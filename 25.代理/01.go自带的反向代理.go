package main

import (
	"github.com/labstack/echo"
	"net/http"
	"net/http/httputil"
)

const (
	URL_HOST = "www.baidu.com"
)
func main() {
	e := echo.New()
	e.GET("/baidu", baidu)
	e.Start(":8080")
}
func baidu(ctx echo.Context) error {
	reverseProxy.ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}
var reverseProxy = &httputil.ReverseProxy{
	Director:directorV250,
}
func directorV250(req *http.Request) {
	req.URL.Scheme = "https"
	req.URL.Host = URL_HOST
	req.Host =URL_HOST
}


