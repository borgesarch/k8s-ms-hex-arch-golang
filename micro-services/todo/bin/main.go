package main

import (
	e "ms-hex-arch-golang-k8s/core/infrastructure/env"
	h "ms-hex-arch-golang-k8s/core/infrastructure/http"
	c "ms-hex-arch-golang-k8s/core/presentation"

	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	m := mux.NewRouter().StrictSlash(true)
	p := h.SSO(1)

	m.HandleFunc("/todos", p.Protect(c.FindAll)).Methods("GET")
	m.HandleFunc("/todos", p.Protect(c.Update)).Methods("PUT")
	m.HandleFunc("/todos", p.Protect(c.Save)).Methods("POST")
	m.HandleFunc("/todos/{id}", p.Protect(c.FindById)).Methods("GET")
	m.HandleFunc("/todos/{id}", p.Protect(c.DeleteById)).Methods("DELETE")

	e.LoadEnv()

	APP_PORT := os.Getenv("APP_PORT")
	http.ListenAndServe(APP_PORT, m)
}
