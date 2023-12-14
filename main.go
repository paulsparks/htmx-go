package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type NavbarProps struct {
	Links []Link
}

type Link struct {
	RouteName string
	URL       string
}

func main() {
	r := mux.NewRouter()

	// page routes
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/test", TestHandler)

	// POST endpoints
	r.HandleFunc("/hello", HelloHandler)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
