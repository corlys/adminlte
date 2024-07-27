package main

import (
	"github.com/corlys/adminlte/app"
)

func main() {
	server := app.SetupApp()
	server.Run()
}
