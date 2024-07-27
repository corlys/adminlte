package controller

import (
	"fmt"
	"net/http"

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
}

func NewUserController(uService service.UserService) UserController {
	return &userController{
		userService: uService,
	}
}

func (c *userController) RenderLogin(ctx *gin.Context) {
	render(ctx, http.StatusOK, views.MakeLoginPage())
	return
}
func (c *userController) RenderRegis(ctx *gin.Context) {
	render(ctx, http.StatusOK, views.MakeRegisterPage())
	return
}
func (c *userController) RenderHome(ctx *gin.Context) {
	render(ctx, http.StatusOK, views.MakeHomePage())
	return
}
func (c *userController) HandleLogin(ctx *gin.Context) {

	var userDto dto.UserLoginRequest
	err := ctx.ShouldBind(&userDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("User Login Failed", err.Error(), http.StatusBadRequest))
	}
	user, err := c.userService.GetUserByEmail(userDto.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("User Login Failed", err.Error(), http.StatusBadRequest))
	}
	fmt.Println(user)
	c.RenderHome(ctx)
}
