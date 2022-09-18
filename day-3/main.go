package main

import (
	"agmc/config"
	"agmc/lib/helpers"
	m "agmc/middlewares"
	"agmc/routes"
)

func init() {
	config.InitDB()

	appEnv := helpers.GetAppEnvConfig()
	if appEnv == "local" {
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
