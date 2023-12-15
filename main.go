package main

import (
	"context"
	"fmt"
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

	conn, err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)

	// page routes
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/test", TestHandler)

	// POST endpoints
	r.HandleFunc("/hello", HelloHandler)
	r.HandleFunc("/todo", TodoHandler)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
