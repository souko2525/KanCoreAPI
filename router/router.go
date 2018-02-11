package router

import (
	"github.com/labstack/echo"
	"github.com/souko2525/KanCoreAPI/api"
)

func Router(e *echo.Echo) {
	v1 := e.Group("api/V1")
	{
		v1.GET("/hello", api.Hello())
		v1.GET("/user/", api.GetUsers())
		v1.GET("/user/:id/", api.GetUser())
	}

}
