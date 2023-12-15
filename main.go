package main

import (
	delete_handlers "htmx-go/handlers/DELETE"
	get_handlers "htmx-go/handlers/GET"
	post_handlers "htmx-go/handlers/POST"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	dbpool, err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	CreateTableIfNotExists(dbpool)

	// page routes
	r.HandleFunc("/", get_handlers.IndexHandler)
	r.HandleFunc("/test", get_handlers.TestHandler)

	// POST endpoints
	r.HandleFunc("/fetchtodo", func(w http.ResponseWriter, r *http.Request) {
		post_handlers.FetchTodoHandler(w, r, dbpool)
	})
	r.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		post_handlers.TodoHandler(w, r, dbpool)
	})

	// DELETE endpoints
	r.HandleFunc("/deletetodo/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		delete_handlers.DeleteTodoHandler(w, r, dbpool)
	})

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
