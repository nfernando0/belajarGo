package main

import (
	"context"
	"fmt"
	"html/template"
	"myblogs/connection"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	ID int
	Title string
	Author string
	Content string
	PostDate time.Time
	FormatDate string
}

func main() {
	e := echo.New();
	connection.DatabaseConnect()


	e.GET("/", home)
	e.GET("/blogs", blogs) // Halaman Blog

	e.POST("/blogs", addBlogs)

	e.Logger.Fatal(e.Start(":5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusMovedPermanently, map[string]string{"Message :": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogs(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, author, content, post_date FROM tb_blog")

	var result []Blog 

	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.Author, &each.Content, &each.PostDate)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		each.FormatDate = each.PostDate.Format("2 januari 2006")
		each.Author = "Fernando"

		result = append(result, each)
	}

	blogs := map[string]interface{} {
		"Blogs" : result,
	}

	var tmpl, err = template.ParseFiles("views/blogs.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), blogs)
}

func addBlogs(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	author := "Fernando"

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog(title, author, content, post_date) VALUES ($1, $2, $3, $4)", title, author, content ,time.Now())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blogs")
}