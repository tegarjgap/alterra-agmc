package controllers

import (
	"myapp/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	storage.InitDB()
	e := echo.New()

	return e
}

func TestGetAllUsers(t *testing.T) {
	var testCase = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "berhasil",
			path:                 "/auth/users",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"message\":\"success get all users\",\"users\":[",
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/auth/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCase {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetAllUsers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}
