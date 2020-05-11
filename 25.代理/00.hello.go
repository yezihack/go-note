package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "hello")
	})
	e.GET("/a", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "hello-aa")
	})
	e.GET("/b", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "hello-bb")
	})
	err := e.StartServer(&http.Server{
		Addr:        ":9091",
		ReadTimeout: 5,
	})
	if err != nil {
		log.Fatal(err)
	}
}
