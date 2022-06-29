package main

import (
	"flag"
	"os"

	"github.com/Final-Project-Kelompok-3/authentications/database"
	"github.com/Final-Project-Kelompok-3/authentications/database/migration"
	"github.com/Final-Project-Kelompok-3/authentications/internal/factory"
	"github.com/Final-Project-Kelompok-3/authentications/internal/http"
	"github.com/Final-Project-Kelompok-3/authentications/internal/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// load env file
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	database.CreateConnection()
	

	var m string // for check migration

	flag.StringVar(
		&m,
		"migrate",
		"run",
		`this argument for check if user want to migrate table, rollback table, or status migration

to use this flag:
	use -migrate=migrate for migrate table
	use -migrate=rollback for rollback table
	use -migrate=status for get status migration`,
	)
	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
		return
	} else if m == "rollback" {
		migration.Rollback()
		return
	} else if m == "status" {
		migration.Status()
		return
	}

	conn := database.GetConnection()
	f := factory.NewFactory(conn)
	e := echo.New()
	http.NewHttp(e, f)
	middleware.Init(e)
	
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}