package main

import (
	"github.com/corlys/adminlte/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func main() {

	app := gin.Default()

	app.Static("/dist", "./dist")

	app.GET("/", func(c *gin.Context) {
		render(c, 200, views.MakeLoginPage())
	})

	app.Run()

}
