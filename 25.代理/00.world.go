package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "world")
	})
	e.GET("/a", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "world-aa")
	})
	e.GET("/b", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "world-bb")
	})
	err := e.StartServer(&http.Server{
		Addr:":9092",
		ReadTimeout:5,
	})
	if err != nil {
		log.Fatal(err)
	}
}
