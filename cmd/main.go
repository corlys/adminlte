package main

import (
	"fmt"
	"time"

	"github.com/corlys/adminlte/app/controller"
	"github.com/corlys/adminlte/app/router"
	"github.com/corlys/adminlte/config"
	"github.com/corlys/adminlte/core/repository"
	"github.com/corlys/adminlte/core/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	db := config.DBSetup()
	defer config.DBClose(db)

	server := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: int(time.Hour.Seconds() * 1),
	})

	server.Use(sessions.Sessions("user", store))
	server.Static("/dist", "./dist")

	newSesionService := service.NewSessionService()

	newUserRepository := repository.NewUserRepository(db)
	newUserService := service.NewUserService(newUserRepository)
	newUserController := controller.NewUserController(newUserService, newSesionService)

	router.UserRouter(server, newUserController)

	err := server.Run()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
