package main

import (
	"fmt"

	"github.com/corlys/adminlte/app/controller"
	"github.com/corlys/adminlte/app/router"
	"github.com/corlys/adminlte/config"
	"github.com/corlys/adminlte/core/repository"
	"github.com/corlys/adminlte/core/service"
	"github.com/gin-gonic/gin"
)

func main() {

	db := config.DBSetup()

	defer config.DBClose(db)

	server := gin.Default()

	server.Static("/dist", "./dist")

	newUserRepository := repository.NewUserRepository(db)
	newUserService := service.NewUserService(newUserRepository)
	newUserController := controller.NewUserController(newUserService)

	router.UserRouter(server, newUserController)

	err := server.Run()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
