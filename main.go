package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/souko2525/KanCoreAPI/database"
	//"github.com/souko2525/KanCoreAPI/models"
	"github.com/souko2525/KanCoreAPI/router"
	//"html/template"
	//"io"
	//"net/http"
)

/*
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "layout.html", "echo server")
}
*/

func main() {
	/*
		t := &Template{
			templates: template.Must(template.ParseFiles("templates/layout.html", "templates/hello.html")),
		}
	*/
	db := database.Connect()
	defer db.Close()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	router.Router(e)
	e.Start(":3000")

	/*
		e.Renderer = t
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "index page")
		})
		e.GET("/hello", Hello)
		e.Logger.Fatal(e.Start(":1323"))
	*/
}
