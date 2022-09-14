package routers

import (
	controller "myapp/dynamic/controllers"

	"github.com/labstack/echo/v4"
)

func GetUserRouter(e *echo.Echo) {
	e.POST("/users", controller.CreateUser)
	e.GET("/users", controller.GetAllUsers)
	e.GET("/users/:id", controller.GetUserById)
	e.PUT("/users/:id", controller.UpdateUserById)
	e.DELETE("/users/:id", controller.DeleteUser)
}
