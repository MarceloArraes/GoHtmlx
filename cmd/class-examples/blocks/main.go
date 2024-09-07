package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Block struct {
	Id int
}

type Blocks struct {
	Start  int
	Next   int
	More   bool
	Blocks []Block
}

type Count struct {
	Count int
}
type Contact struct {
	Name  string
	Email string
}

func newContact(name string, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: []Contact{
			newContact("John", "jd@gmail.com"),
			newContact("Clara", "cd@gmail.com"),
		},
	}
}

func main() {
	e := echo.New()
	// e.Renderer = newTemplates()

	e.Use(middleware.Logger())

	// count := Count{Count: 0}
	data := newData()
	e.Renderer = NewTemplate()

	e.GET("/", func(c echo.Context) error {
		// console.log('what');
		return c.Render(200, "index", data)

	})

	// e.POST("/count", func(c echo.Context) error {
	// 	count.Count++
	// 	return c.Render(200, "count", count)
	// })
	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		data.Contacts = append(data.Contacts, newContact(name, email))
		return c.Render(200, "display", data)
	})

	/* e.GET("/blocks", func(c echo.Context) error {
		startStr := c.QueryParam("start")
		start, err := strconv.Atoi(startStr)
		if err != nil {
			start = 0
		}

		blocks := []Block{}
		for i := start; i < start+10; i++ {
			blocks = append(blocks, Block{Id: i})
		}

		template := "blocks"
		if start == 0 {
			template = "blocks-index"
		}
		return c.Render(http.StatusOK, template, Blocks{
			Start:  start,
			Next:   start + 10,
			More:   start+10 < 100,
			Blocks: blocks,
		})
	}) */

	e.Logger.Fatal(e.Start(":42069"))
}
