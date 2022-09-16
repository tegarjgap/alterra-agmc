package main

import (
	"myapp/routers"
	"myapp/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	storage.InitDB()
	storage.InitialMigration()

	routers.GetUserRouter(e)

	e.Logger.Fatal(e.Start(":14045"))
}
