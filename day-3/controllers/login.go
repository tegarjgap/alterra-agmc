package controllers

import (
	"myapp/lib/database"
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserLogin(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)

	users, err := database.UserLogin(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

func Booklogin(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)

	users, err := database.BookLogin(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login for book",
		"users":  users,
	})

}
