package router

import (
	"github.com/corlys/adminlte/app/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, userC controller.UserController) {
	userRoutes := router.Group("/")
	{
		userRoutes.GET("", userC.RenderHome)
		userRoutes.GET("register", userC.RenderRegis)
		userRoutes.GET("login", userC.RenderLogin)
		userRoutes.GET("totp-setup", userC.RenderTotpSetup)
		userRoutes.GET("totp-verify", userC.RenderTotpVerify)
		userRoutes.GET("logout", userC.HandleLogout)

		userRoutes.POST("login", userC.HandleLogin)
		userRoutes.POST("register", userC.HandleRegis)
		userRoutes.POST("totp-setup", userC.HandleTotpSetup)
	}
}
