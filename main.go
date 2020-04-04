package main

import (
	"github.com/datt16/Foooo/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Logger())
	e.GET("/hello", handler.MainPage())
	e.Start(":8081")
}
