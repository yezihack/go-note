package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "net/http/pprof"
)
func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	r := gin.Default()
	r.Run(":8080")
}


