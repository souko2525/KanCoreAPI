package api

import (
	"net/http"

	"github.com/labstack/echo"
)

//Hello return echo.HandlerFunc
func Hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
