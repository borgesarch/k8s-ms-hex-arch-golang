package auth

import (
	s "clean-code-golang/core/infrastructure/ioc/services"

	"net/http"
)

func Handle(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		c := s.Resolver()

		c.Invoke(func(config *s.Config) {

			tokenString := r.Header.Get("Authorization")

			if tokenString == "" {
				http.Error(w, "Header inválido!", http.StatusUnauthorized)
			}

			token, err := config.AuthService.CheckToken(tokenString)

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}

			if !token {
				http.Error(w, "Não autorizado!", http.StatusUnauthorized)
			}

			f(w, r)

		})
	}
}
