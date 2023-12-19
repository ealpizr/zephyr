package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RootResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func RootHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, &RootResponse{
		Status: http.StatusOK,
		Message: "Hello from Zephyr!",
	})
}