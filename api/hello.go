package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
