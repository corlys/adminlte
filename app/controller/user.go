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
	userService    service.UserService
	sessionService service.SessionService
}

type UserController interface {
	RenderLogin(ctx *gin.Context)
	RenderRegis(ctx *gin.Context)
	RenderHome(ctx *gin.Context)
	HandleLogin(ctx *gin.Context)
	HandleRegis(ctx *gin.Context)
	HandleLogout(ctx *gin.Context)
}

func NewUserController(uService service.UserService, sService service.SessionService) UserController {
	return &userController{
		userService:    uService,
		sessionService: sService,
	}
}

func (c *userController) RenderLogin(ctx *gin.Context) {
	render(ctx, http.StatusOK, views.MakeLoginPage())
}
func (c *userController) RenderRegis(ctx *gin.Context) {
	render(ctx, http.StatusOK, views.MakeRegisterPage())
}
func (c *userController) RenderHome(ctx *gin.Context) {
	user := c.sessionService.GetUserSession(ctx)
	fmt.Println(user)
	render(ctx, http.StatusOK, views.MakeHomePage(user))
}
func (c *userController) HandleLogin(ctx *gin.Context) {
	var userDto dto.UserLoginRequest
	err := ctx.ShouldBind(&userDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("User Login Failed", err.Error(), http.StatusBadRequest))
	}
	loggedIn := c.userService.VerifyLogin(userDto.Email, userDto.Password)
	if !loggedIn {
		render(ctx, http.StatusOK, views.MakeLoginPage())
		return
	}
	user, _ := c.userService.GetUserByEmail(userDto.Email)
	c.sessionService.SetUserSession(ctx, user)
	ctx.Redirect(http.StatusSeeOther, "/")
}
func (c *userController) HandleRegis(ctx *gin.Context) {
	var userDto dto.UserRegisterRequest
	err := ctx.ShouldBind(&userDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("User Regis Failed", err.Error(), http.StatusBadRequest))
	}
	user, erro := c.userService.RegisterUser(userDto)
	if erro != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("User Regis Failed", err.Error(), http.StatusBadRequest))
	}
	c.sessionService.SetUserSession(ctx, user)
	ctx.Redirect(http.StatusSeeOther, "/")
}
func (c *userController) HandleLogout(ctx *gin.Context) {
	c.sessionService.DeleteUserSession(ctx)
	ctx.Redirect(http.StatusSeeOther, "/")
}
