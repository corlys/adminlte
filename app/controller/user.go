package controller

import (
	"net/http"
	"fmt"

	"github.com/corlys/adminlte/common/base"
	"github.com/corlys/adminlte/core/helper/dto"
	"github.com/corlys/adminlte/core/service"
	"github.com/corlys/adminlte/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

type userController struct {
	userService service.UserService
}

type UserController interface {
	RenderLogin(ctx *gin.Context)
	RenderRegis(ctx *gin.Context)
	RenderHome(ctx *gin.Context)
	HandleLogin(ctx *gin.Context)
	HandleRegis(ctx *gin.Context)
}

func NewUserController(uService service.UserService) UserController {
	return &userController{
		userService: uService,
	}
}

func (c *userController) RenderLogin(ctx *gin.Context) {
	render(ctx, http.StatusOK, views.MakeLoginPage())
}
func (c *userController) RenderRegis(ctx *gin.Context) {
	render(ctx, http.StatusOK, views.MakeRegisterPage())
}
func (c *userController) RenderHome(ctx *gin.Context) {
	render(ctx, http.StatusOK, views.MakeHomePage())
}
func (c *userController) HandleLogin(ctx *gin.Context) {
	var userDto dto.UserLoginRequest
	err := ctx.ShouldBind(&userDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("User Login Failed", err.Error(), http.StatusBadRequest))
	}
	loggedIn := c.userService.VerifyLogin(userDto.Email, userDto.Password)
	fmt.Println(loggedIn, userDto)
	if loggedIn {
		render(ctx, http.StatusOK, views.MakeHomePage())		
		return
	} else {
		render(ctx, http.StatusOK, views.MakeLoginPage())
		return
	}
}
func (c *userController) HandleRegis(ctx *gin.Context) {
	var userDto dto.UserRegisterRequest
	err := ctx.ShouldBind(&userDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("User Regis Failed", err.Error(), http.StatusBadRequest))
	}
	c.userService.RegisterUser(userDto)
	render(ctx, http.StatusOK, views.MakeHomePage())
}
