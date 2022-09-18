package middlewares

import (
	"myapp/constanta"
	"myapp/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateTokenUsingCustomClaims(userName string) (string, error) {
	// create custom claims first
	claims := &models.CustomClaims{
		UserName: userName,
		Admin:    true,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// and then create token from custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// return token
	return token.SignedString([]byte(constanta.JWT_SECRET))

}
