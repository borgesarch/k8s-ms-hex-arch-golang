package main

import (
	e "clean-code-golang/core/infrastructure/env"
	c "clean-code-golang/core/presentation"
	a "clean-code-golang/core/presentation/auth"

	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter().StrictSlash(true)

	m.HandleFunc("/users", a.Handle(c.FindAll)).Methods("GET")
	m.HandleFunc("/users", a.Handle(c.Update)).Methods("PUT")
	m.HandleFunc("/users", a.Handle(c.Save)).Methods("POST")
	m.HandleFunc("/users/{id}", a.Handle(c.FindById)).Methods("GET")
	m.HandleFunc("/users/{id}", a.Handle(c.DeleteById)).Methods("DELETE")

	m.HandleFunc("/auth", a.Auth).Methods("POST")
	m.HandleFunc("/sign", a.Sign).Methods("POST")

	e.LoadEnv()

	APP_PORT := os.Getenv("APP_PORT")
	http.ListenAndServe(APP_PORT, m)
}
