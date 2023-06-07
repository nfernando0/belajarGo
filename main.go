package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Title string
	Content string
	Author string
	postDate string
}

var dataProjects = []Project {
	{
		Title : "Hello",
	Content: "Ini Content",
	Author: "Fernando",
	postDate: "10/10/2023",
	},
	{
		Title : "Hello1",
	Content: "Ini Content 1",
	Author: "Fernando",
	postDate: "10/10/2023",
	},
}

func main() {
	e := echo.New()

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World")
	// })

	e.Static("/public", "public")
	// Routing
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/projects", projects) // projects
	e.GET("/project/:id", projectDetail) // detail project
	e.GET("/formAddProjects", formAddProjects) // tambah project
	e.GET("/testimonial", testimonial)
	
	e.POST("/addProjects", addProjects)
	// e.POST("/deleteProject/:id", deleteProject) // Hapus Project
	e.Logger.Fatal(e.Start("localhost:5000"))

}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}
func projects(c echo.Context) error {

	var tmpl, err = template.ParseFiles("views/projects.html")
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{} {
		"Projects": dataProjects,
	}

	return tmpl.Execute(c.Response(), projects)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// data := map[string]interface{} {
	// 	"id" : id,
	// 	"title" : "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
	// 	"content" : "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian Manpower Group, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Khusus di sektor teknologi yang berkembang pesat, menurut Kemendikbudristek, Indonesia kekurangan sembilan juta pekerja teknologi hingga tahun 2030. Hal itu berarti Indonesia memerlukan sekitar 600 ribu SDM digital yang memasuki pasar setiap tahunnya.",
	// }

	var projectDetail = Project{}

	for i, data := range dataProjects {
		if id == i {
			projectDetail = Project{
				Title: data.Title,
				Content: data.Content,
				postDate: data.postDate,
				Author: data.Author,
			}
		}
	}

	data := map[string]interface{} {
		"Projects" : projectDetail,
	}

	var tmpl, err = template.ParseFiles("views/project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}


func formAddProjects (c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-projects.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

// nambah project
func addProjects(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("desc")

	println("title : " + title)
	println("content : " + content)

	return c.Redirect(http.StatusMovedPermanently, "/projects")
}

// func projectDetail(c echo.Context) error {
// 	data := map[string]interface{} {
// 		"login" : true
// 	}

// 	var tmpl,err = template.ParseFiles("views")
// }

func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message" : err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}


