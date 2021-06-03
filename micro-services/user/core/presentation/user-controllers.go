package presentation

import (
	t "clean-code-golang/core/domain/users"
	i "clean-code-golang/core/infrastructure/ioc/repositories"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func FindAll(w http.ResponseWriter, r *http.Request) {

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		json.NewEncoder(w).Encode(config.UserRepository.FindAll())
	})

}

func FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		json.NewEncoder(w).Encode(config.UserRepository.FindById(id))
	})

}

func DeleteById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		config.UserRepository.DeleteById(id)
	})

	json.NewEncoder(w).Encode(true)
}

func Save(w http.ResponseWriter, r *http.Request) {

	user := t.User{}

	json.NewDecoder(r.Body).Decode(&user)

	c := i.Resolver()

	err := c.Invoke(func(config *i.Config) {
		config.UserRepository.Save(user)
	})

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(&user)
}

func Update(w http.ResponseWriter, r *http.Request) {

	user := t.User{}

	json.NewDecoder(r.Body).Decode(&user)

	c := i.Resolver()

	err := c.Invoke(func(config *i.Config) {
		config.UserRepository.Update(user)
	})

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(&user)
}
