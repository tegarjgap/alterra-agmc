package main

import (
	"myapp/dynamic/routers"
	"myapp/dynamic/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	storage.InitDB()
	storage.InitialMigration()

	routers.GetUserRouter(e)

	e.Logger.Fatal(e.Start(":14045"))
}
