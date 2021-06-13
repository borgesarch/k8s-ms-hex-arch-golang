package http

import (
	"time"

	"net/http"

	"github.com/MicahParks/keyfunc"

	jwt "github.com/dgrijalva/jwt-go"
)

type ISSO interface {
	Protect(f http.HandlerFunc) http.HandlerFunc
}

type SSO int

func (s SSO) Protect(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		jwksUrl := APP_PORT := os.Getenv("OPENID_CERT")

		refreshInterval := time.Hour
		options := keyfunc.Options{
			RefreshInterval: &refreshInterval,
			RefreshErrorHandler: func(err error) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			},
		}

		jwks, err := keyfunc.Get(jwksUrl, options)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, jwks.KeyFunc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Token inválido!", http.StatusUnauthorized)
			return
		}

		f(w, r)

	}
}
