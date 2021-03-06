package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/souko2525/KanCoreAPI/database"
	"github.com/souko2525/KanCoreAPI/router"
	//"html/template"
	//"io"
	//"net/http"
)

func main() {
	if err := database.Connect(); err != nil {
		os.Exit(1)
	}
	defer database.Close()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	router.Router(e)
	e.Logger.Fatal(e.Start("127.0.0.1:3000"))

}
