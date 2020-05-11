package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	e := echo.New()
	e.GET("/v", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "v100")
	})
	//用于监控收集数据的端口
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	e.Logger.Fatal(e.Start(":8080"))
}
