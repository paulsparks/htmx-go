package main

import (
	deletehandler "htmx-go/handlers/DELETE"
	gethandler "htmx-go/handlers/GET"
	posthandler "htmx-go/handlers/POST"
	puthandler "htmx-go/handlers/PUT"
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

	// GET endpoints
	r.HandleFunc("/", gethandler.IndexHandler)
	r.HandleFunc("/updatetodoform/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		gethandler.UpdateTodoFormHandler(w, r, dbpool)
	})
	r.HandleFunc("/test", gethandler.TestHandler)

	// POST endpoints
	r.HandleFunc("/fetchtodo", func(w http.ResponseWriter, r *http.Request) {
		posthandler.FetchTodoHandler(w, r, dbpool)
	})
	r.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		posthandler.TodoHandler(w, r, dbpool)
	})

	// DELETE endpoints
	r.HandleFunc("/deletetodo/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		deletehandler.DeleteTodoHandler(w, r, dbpool)
	})

	// PUT endpoints
	r.HandleFunc("/updatetodo/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		puthandler.UpdateTodoHandler(w, r, dbpool)
	})

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
