package presentation

import (
	t "clean-code-golang/core/domain/todo"
	i "clean-code-golang/core/infrastructure/ioc/repositories"
	s "clean-code-golang/core/infrastructure/ioc/services"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func FindAll(w http.ResponseWriter, r *http.Request) {

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		json.NewEncoder(w).Encode(config.TodoRepository.FindAll())
	})

}

func FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		json.NewEncoder(w).Encode(config.TodoRepository.FindById(id))
	})

}

func DeleteById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	c := i.Resolver()

	c.Invoke(func(config *i.Config) {
		config.TodoRepository.DeleteById(id)
	})

	json.NewEncoder(w).Encode(true)
}

func Save(w http.ResponseWriter, r *http.Request) {

	todo := t.Todo{}

	json.NewDecoder(r.Body).Decode(&todo)

	c := s.Resolver()

	err := c.Invoke(func(config *s.Config) {
		config.TodoService.Save(todo)
	})

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(&todo)
}

func Update(w http.ResponseWriter, r *http.Request) {

	todo := t.Todo{}

	json.NewDecoder(r.Body).Decode(&todo)

	c := s.Resolver()

	err := c.Invoke(func(config *s.Config) {
		config.TodoService.Update(todo)
	})

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(&todo)
}
