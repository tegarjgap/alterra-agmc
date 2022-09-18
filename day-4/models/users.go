package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type CustomClaims struct {
	UserName string `json:"user_name"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}

type Users struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token"`
}

type Books struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	AmountPage int    `json:"amount_page"`
}
