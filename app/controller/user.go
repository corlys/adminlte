package controller

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/corlys/adminlte/common/base"
	"github.com/corlys/adminlte/core/helper/dto"
	"github.com/corlys/adminlte/core/service"
	"github.com/corlys/adminlte/views"
	"github.com/pquerna/otp"

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
	RenderTotpSetup(ctx *gin.Context)
	RenderTotpVerify(ctx *gin.Context)
	HandleLogin(ctx *gin.Context)
	HandleRegis(ctx *gin.Context)
	HandleLogout(ctx *gin.Context)
	HandleTotpSetup(ctx *gin.Context)
	HandleTotpVerify(ctx *gin.Context)
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
func (c *userController) RenderTotpSetup(ctx *gin.Context) {
	keyUrl := c.sessionService.GetTotpKeySession(ctx)
	key, err := otp.NewKeyFromURL(keyUrl)
	if err != nil {
		fmt.Println("Error : ", err)
		ctx.Redirect(http.StatusSeeOther, "/register")
		return
	}
	qrCodeUrl := fmt.Sprintf(
		"https://image-charts.com/chart?chs=200x200&cht=qr&chl=%s&choe=UTF-8",
		url.QueryEscape(key.URL()),
	)
	render(ctx, http.StatusOK, views.MakeTOTPSetupPage(qrCodeUrl, key.AccountName()))
}
func (c *userController) RenderTotpVerify(ctx *gin.Context) {
	keyUrl := c.sessionService.GetTotpKeySession(ctx)
	key, err := otp.NewKeyFromURL(keyUrl)
	if err != nil {
		fmt.Println("Error : ", err)
		ctx.Redirect(http.StatusSeeOther, "/login")
		return
	}
	render(ctx, http.StatusOK, views.MakeTOTPVerifyPage(key.AccountName()))
}
func (c *userController) HandleLogin(ctx *gin.Context) {
	var userDto dto.UserLoginRequest
	err := ctx.ShouldBind(&userDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("User Login Failed", err.Error(), http.StatusBadRequest))
	}
	fmt.Println("UserLoginRequest: ", userDto)
	loggedIn := c.userService.VerifyLogin(userDto.Email, userDto.Password)
	if !loggedIn {
		fmt.Println("wrong pass", userDto)
		render(ctx, http.StatusOK, views.MakeLoginPage())
		return
	}
	user, _ := c.userService.GetUserByEmail(userDto.Email)
	url := c.userService.GetUserTotp(user.Email)
	c.sessionService.SetTotpKeySession(ctx, url)
	ctx.Redirect(http.StatusSeeOther, "/totp-verify")
}
func (c *userController) HandleRegis(ctx *gin.Context) {
	var userDto dto.UserRegisterRequest
	err := ctx.ShouldBind(&userDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("User Regis Failed", err.Error(), http.StatusBadRequest))
	}
	fmt.Println("UserRegisterRequest: ", userDto)
	user, erro := c.userService.RegisterUser(userDto)
	if erro != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("User Regis Failed", err.Error(), http.StatusBadRequest))
	}
	secret, err := c.userService.GenerateTotp(user.Email)
	c.sessionService.SetTotpKeySession(ctx, secret.URL())
	ctx.Redirect(http.StatusSeeOther, "/totp-setup")
}
func (c *userController) HandleLogout(ctx *gin.Context) {
	c.sessionService.DeleteUserSession(ctx)
	c.sessionService.DeleteTotpKeySession(ctx)
	ctx.Redirect(http.StatusSeeOther, "/")
}
func (c *userController) HandleTotpSetup(ctx *gin.Context) {
	var totpDto dto.UserTotpRequest
	err := ctx.ShouldBind(&totpDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("Failed to send TOTP", err.Error(), http.StatusBadRequest))
	}
	res := c.userService.ValidateTotp(totpDto.AccountName, totpDto.OtpCode)
	if !res {
		fmt.Println("Totp Not Valid")
		ctx.Redirect(http.StatusSeeOther, "/totp-setup")
	} else {
		user, err := c.userService.GetUserByEmail(totpDto.AccountName)
		if err != nil {
			fmt.Println("User fetching failed")
			ctx.Redirect(http.StatusSeeOther, "/totp-setup")
		}
		c.sessionService.SetUserSession(ctx, user)
		ctx.Redirect(http.StatusSeeOther, "/")
	}
}
func (c *userController) HandleTotpVerify(ctx *gin.Context) {
	var totpDto dto.UserTotpRequest
	err := ctx.ShouldBind(&totpDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, base.CreateFailResponse("Failed to send TOTP", err.Error(), http.StatusBadRequest))
	}
	res := c.userService.ValidateTotp(totpDto.AccountName, totpDto.OtpCode)
	if !res {
		fmt.Println("Totp Not Valid")
		ctx.Redirect(http.StatusSeeOther, "/totp-verify")
	} else {
		user, err := c.userService.GetUserByEmail(totpDto.AccountName)
		if err != nil {
			fmt.Println("User fetching failed")
			ctx.Redirect(http.StatusSeeOther, "/totp-verify")
		}
		c.sessionService.SetUserSession(ctx, user)
		ctx.Redirect(http.StatusSeeOther, "/")
	}
}
