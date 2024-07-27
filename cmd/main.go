package main

import (
	"github.com/corlys/adminlte/app"
	"github.com/corlys/adminlte/config"
)

func main() {

	db := config.DBSetup()

	defer config.DBClose(db)

	server := app.SetupApp(db)
	server.Run()
}
