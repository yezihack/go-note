package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http/httputil"
	"net/url"
)

// 代理server0, server1
var URLs = []*url.URL{
	{
		Scheme: "http",
		Host:   "localhost:7000",
	},
	{
		Scheme: "http",
		Host:   "localhost:7001",
	},
}

func NewProxy(c *gin.Context) {
	target := URLs[rand.Int() % len(URLs)]
	fmt.Printf("host:%s\n", target.Host)
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ServeHTTP(c.Writer, c.Request)
}
func AddRouter(router *gin.Engine) {
	router = gin.Default()
	router.GET("/ping", NewProxy)
	router.POST("/user",NewProxy)
}
func main() {
	router := gin.Default()
	AddRouter(router)
	router.Run(":8080")
}
