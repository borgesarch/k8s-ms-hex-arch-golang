package auth

import (
	t "clean-code-golang/core/domain/users"
	n "clean-code-golang/core/infrastructure/ioc/commons"
	i "clean-code-golang/core/infrastructure/ioc/repositories"
	w "clean-code-golang/core/infrastructure/jwt"

	"errors"
	"time"

	"os"

	"github.com/dgrijalva/jwt-go"
)

type IAuthService interface {
	Auth(user t.User) (string, error)
	Sign(user t.User) t.User
	CheckToken(tokenString string) (bool, error)
}

type AuthService int

func (o AuthService) Sign(user t.User) (t.User, error) {

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {

		x := n.Resolver()

		x.Invoke(func(con *n.Config) {
			user.Password = con.EncrityHash.HashPassword(user.Password)
		})

		config.UserRepository.Save(user)

	})

	return user, nil
}

func (o AuthService) CheckToken(tokenString string) (bool, error) {

	APP_SIGNING_KEY := os.Getenv("APP_SIGNING_KEY")
	signingKey := []byte(APP_SIGNING_KEY)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	return token.Valid, err
}

func (o AuthService) Auth(user t.User) (string, error) {

	c := i.Resolver()

	err := c.Invoke(func(config *i.Config) {

		userFound, err := config.UserRepository.FindByEmail(user.Email)

		if err != nil {
			err = errors.New("User not found!")
		}

		x := n.Resolver()

		doMach := x.Invoke(func(con *n.Config) bool {
			return con.EncrityHash.DoPasswordsMatch(
				con.EncrityHash.HashPassword(user.Password), userFound.Password)
		})

		if doMach != nil {
			err = errors.New("Password do not match!")
		}
	})

	if err != nil {
		err = errors.New(err.Error())
	}

	claims := w.CustomClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(15, 1, 1).UTC().Unix(),
			Issuer:    "core:go",
		},
	}

	APP_SIGNING_KEY := os.Getenv("APP_SIGNING_KEY")
	signingKey := []byte(APP_SIGNING_KEY)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	return ss, err
}
