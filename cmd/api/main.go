package main

import (
	"github.com/ealpizr/zephyr/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true

  e.GET("/", handlers.RootHandler)

  e.Logger.Fatal(e.Start(":3000"))
}