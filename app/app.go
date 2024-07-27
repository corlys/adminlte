package app

import (
	"fmt"
	"net/http"

	"github.com/corlys/adminlte/common"
	"github.com/corlys/adminlte/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterRequest struct {
	LoginRequest
	FullName string `form:"full_name" binding:"required"`
}

func SetupApp() *gin.Engine {

	app := gin.Default()

	app.Static("/dist", "./dist")

	app.GET("/", func(c *gin.Context) {
		render(c, 200, views.MakeHomePage())
	})

	app.GET("/register", func(c *gin.Context) {
		render(c, 200, views.MakeRegisterPage())
	})

	app.GET("/login", func(c *gin.Context) {
		render(c, 200, views.MakeLoginPage())
	})

	app.POST("/login", func(c *gin.Context) {

		var loginRequest LoginRequest

		if err := c.ShouldBind(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, common.MsgErrorRes("bad request"))
			return
		}

		fmt.Println(loginRequest)

	})

	app.POST("/register", func(c *gin.Context) {

		var registerRequest RegisterRequest

		if err := c.ShouldBind(&registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, common.MsgErrorRes("bad request"))
			return
		}

		fmt.Println(registerRequest)

	})

	return app

}
