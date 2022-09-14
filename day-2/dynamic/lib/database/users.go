package database

import (
	"myapp/dynamic/models"
	"myapp/dynamic/storage"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) (interface{}, error) {
	user := models.Users{}
	c.Bind(&user)

	if err := storage.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetAllUsers() (interface{}, error) {
	var users []models.Users

	if err := storage.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(c echo.Context) (interface{}, error) {
	var user models.Users
	userID, _ := strconv.Atoi(c.Param("id"))

	if err := storage.DB.Where("ID = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserById(c echo.Context) (interface{}, error) {
	user := models.Users{}
	c.Bind(&user)

	userID, _ := strconv.Atoi(c.Param("id"))

	if err := storage.DB.Where("ID = ?", userID).Updates(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUserById(c echo.Context) (interface{}, error) {
	var user models.Users
	userID, _ := strconv.Atoi(c.Param("id"))

	if err := storage.DB.Where("ID = ?", userID).Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
