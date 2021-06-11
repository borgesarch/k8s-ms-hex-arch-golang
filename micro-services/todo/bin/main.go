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

	m.HandleFunc("/todos", c.FindAll).Methods("GET")
	m.HandleFunc("/todos", c.Update).Methods("PUT")
	m.HandleFunc("/todos", c.Save).Methods("POST")
	m.HandleFunc("/todos/{id}", c.FindById).Methods("GET")
	m.HandleFunc("/todos/{id}", c.DeleteById).Methods("DELETE")

	e.LoadEnv()

	APP_PORT := os.Getenv("APP_PORT")
	http.ListenAndServe(APP_PORT, m)
}
