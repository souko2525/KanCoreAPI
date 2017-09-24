package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "content", "world")
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseFiles("templates/layout.html", "templates/hello.html")),
		//templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "index page")
	})
	e.GET("/hello", Hello)
	e.Logger.Fatal(e.Start(":1323"))
}
