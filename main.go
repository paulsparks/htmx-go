package main

import (
	"log"
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

	dbpool, err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	CreateTableIfNotExists(dbpool)

	// page routes
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/test", TestHandler)

	// POST endpoints
	r.HandleFunc("/fetchtodo", func(w http.ResponseWriter, r *http.Request) {
		FetchTodoHandler(w, r, dbpool)
	})
	r.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		TodoHandler(w, r, dbpool)
	})

	// DELETE endpoints
	r.HandleFunc("/deletetodo/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		DeleteTodoHandler(w, r, dbpool)
	})

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
