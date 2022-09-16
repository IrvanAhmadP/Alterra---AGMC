package main

import (
	"agmc/config"
	m "agmc/middlewares"
	"agmc/routes"
)

func init() {
	config.InitDB()

	configENV := config.GetConfig()
	if configENV["APP_ENV"] == "local" {
		config.InitMigrate()
	}

}

func main() {
	e := routes.New()

	// log middleware
	m.LogMiddleware(e)

	//start the server
	e.Logger.Fatal(e.Start(":8000"))
}
