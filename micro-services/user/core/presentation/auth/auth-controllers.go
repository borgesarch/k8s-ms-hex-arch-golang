package auth

import (
	t "clean-code-golang/core/domain/users"

	s "clean-code-golang/core/infrastructure/ioc/services"

	"encoding/json"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	user := t.User{}

	c := s.Resolver()

	json.NewDecoder(r.Body).Decode(&user)

	c.Invoke(func(config *s.Config) {
		token, err := config.AuthService.Auth(user)

		if err != nil {
			http.Error(w, err.Error(), 400)
		}

		json.NewEncoder(w).Encode(token)
	})

}

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

func Sign(w http.ResponseWriter, r *http.Request) {
	user := t.User{}

	json.NewDecoder(r.Body).Decode(&user)

	c := s.Resolver()

	c.Invoke(func(config *s.Config) {

		user, err := config.AuthService.Sign(user)

		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}

		json.NewEncoder(w).Encode(user)

	})
}
