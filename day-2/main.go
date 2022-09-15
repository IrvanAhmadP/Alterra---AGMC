package main

import (
	"agmc/config"
	"agmc/routes"
	"fmt"
)

func init() {
	config.InitDB()

	configENV := config.GetConfig()
	if configENV["APP_ENV"] == "local" {
		config.InitMigrate()
	}
	fmt.Println(configENV)
}

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
