package main

import (
	"github.com/labstack/echo"
	"fmt"
)

func main() {
	e := echo.New()
	fmt.Println(e)
}
