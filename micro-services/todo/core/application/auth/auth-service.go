package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type IAuthService interface {
	CheckToken(tokenString string) (bool, error)
}

type AuthService int

func (o AuthService) CheckToken(tokenString string) (bool, error) {

	APP_SIGNING_KEY := os.Getenv("APP_SIGNING_KEY")
	signingKey := []byte(APP_SIGNING_KEY)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	return token.Valid, err
}
