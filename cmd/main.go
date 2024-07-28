package main

import (
	"encoding/gob"
	"fmt"
	"time"

	"github.com/corlys/adminlte/app/controller"
	"github.com/corlys/adminlte/app/router"
	"github.com/corlys/adminlte/config"
	"github.com/corlys/adminlte/core/helper/dto"
	"github.com/corlys/adminlte/core/repository"
	"github.com/corlys/adminlte/core/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading environment table")
		panic(err)
	}

	db := config.DBSetup()
	defer config.DBClose(db)

	server := gin.Default()

	gob.Register(dto.UserResponse{})

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

	errServ := server.Run()
	if errServ != nil {
		fmt.Println(err)
		panic(errServ)
	}

}
