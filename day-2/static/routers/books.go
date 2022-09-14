package routers

import (
	"myapp/static/controllers"

	"github.com/labstack/echo/v4"
)

func GetRouterBooks() {
	e := echo.New()

	e.POST("/books", controllers.CreateBook)
	e.GET("/books", controllers.GetAllBooks)
	e.GET("/books/:id", controllers.GetBookById)
	e.PUT("/books/:id", controllers.UpdateBookById)
	e.DELETE("/books/:id", controllers.DeleteBookById)

	e.Logger.Fatal(e.Start(":3000"))
}
