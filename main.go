package main

import (
	"html/template"
	"myblogs/connection"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New();
	connection.DatabaseConnect()


	e.GET("/", home)
	e.GET("/blogs", blog) // Halaman Blog

	e.Logger.Fatal(e.Start(":5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusMovedPermanently, map[string]string{"Message :": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/blogs.html")

	if err != nil {
		return c.JSON(http.StatusMovedPermanently, map[string]string{"Message :": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}