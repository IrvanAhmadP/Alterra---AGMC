package main

import (
	"agmc-day-6/database"
	"agmc-day-6/database/migration"
	"agmc-day-6/database/seeder"
	"agmc-day-6/internal/factory"
	"agmc-day-6/internal/http"
	"flag"
	"fmt"
	"os"

	"agmc-day-6/internal/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// load env configuration
func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	database.CreateConnection()

	var m string // for check migration
	var s string // for check migration

	flag.StringVar(
		&m,
		"m",
		"none",
		`this argument for check if user want to migrate table, rollback table, or status migration

to use this flag:
	use -m=migrate for migrate table
	use -m=rollback for rollback table
	use -m=status for get status migration`,
	)
	flag.Parse()

	flag.StringVar(
		&s,
		"s",
		"none",
		`this argument for check if user want to seed table
to use this flag:
	use -s=all to seed all table`,
	)

	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
	} else if m == "rollback" {
		migration.Rollback()
	} else if m == "status" {
		migration.Status()
	}

	if s == "seed" {
		seeder.NewSeeder().SeedAll()
	} else if s == "fresh" {
		seeder.NewSeeder().DeleteAll()
		seeder.NewSeeder().SeedAll()
	}

	f := factory.NewFactory()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e, f)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
