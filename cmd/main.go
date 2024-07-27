package main

import (
	"github.com/corlys/adminlte/app"
	"github.com/corlys/adminlte/config"
)

func main() {

	db := config.DBSetup()

	server := app.SetupApp(db)
	server.Run()
}
