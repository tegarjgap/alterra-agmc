package database

import (
	"myapp/middlewares"
	"myapp/models"
	"myapp/storage"

	"github.com/labstack/echo/v4"
)

func UserLogin(user *models.Users) (interface{}, error) {
	var err error

	if err := storage.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateTokenUsingCustomClaims(user.Name)
	if err != nil {
		return nil, err
	}

	if err := storage.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func BookLogin(c echo.Context) (interface{}, error) {
	var err error

	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "root" || password != "huft" {
		return nil, echo.ErrUnauthorized
	}

	token, err := middlewares.CreateTokenUsingCustomClaims(username)
	if err != nil {
		return nil, err
	}

	return token, err
}
