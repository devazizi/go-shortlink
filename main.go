package main

import (
	"goshorter/adapter"
	"goshorter/router"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func loadEnvFile() {
	godotenv.Load(".env")
}

func main() {
	loadEnvFile()
	db := adapter.NewDB(os.Getenv("DB_DSN"))
	e := echo.New()
	router.RegisterRouteList(e, db)
	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}
