package routers

import (
	"myapp/constanta"
	"myapp/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetUserRouter(e *echo.Echo) {
	e.POST("/login", controllers.UserLogin)
	e.POST("/login/book", controllers.Booklogin)

	// not authenticated user & book
	e.POST("/users", controllers.CreateUser)
	e.GET("/books", controllers.GetAllBook)
	e.GET("/books/:id", controllers.GetBookById)

	r := e.Group("/auth")
	r.Use(middleware.JWT([]byte(constanta.JWT_SECRET)))
	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// authenticated user
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserById)
	r.PUT("/users/:id", controllers.UpdateUserById)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// authenticated book
	r.POST("/books", controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBookById)
	r.DELETE("/books/:id", controllers.DeleteBookById)
}
