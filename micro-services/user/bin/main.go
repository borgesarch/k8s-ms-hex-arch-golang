package main

import (
	e "ms-hex-arch-golang-k8s/core/infrastructure/env"
	c "ms-hex-arch-golang-k8s/core/presentation"

	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter().StrictSlash(true)

	m.HandleFunc("/users", c.FindAll).Methods("GET")
	m.HandleFunc("/users", c.Update).Methods("PUT")
	m.HandleFunc("/users", c.Save).Methods("POST")
	m.HandleFunc("/users/{id}", c.FindById).Methods("GET")
	m.HandleFunc("/users/{id}", c.DeleteById).Methods("DELETE")
	e.LoadEnv()

	APP_PORT := os.Getenv("APP_PORT")
	http.ListenAndServe(APP_PORT, m)
}
