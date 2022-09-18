package controllers

import (
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	books = map[int]*models.Books{}
	seq   = 1
)

func CreateBook(c echo.Context) error {
	book := &models.Books{
		ID: seq,
	}

	if err := c.Bind(book); err != nil {
		return err
	}

	books[book.ID] = book
	seq++
	return c.JSON(http.StatusCreated, book)
}

func GetAllBook(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, books[id])
}

func UpdateBookById(c echo.Context) error {
	book := new(models.Books)
	if err := c.Bind(book); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Title = book.Title
	return c.JSON(http.StatusOK, books[id])
}

func DeleteBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	return c.JSON(http.StatusOK, books)
}
