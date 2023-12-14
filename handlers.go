package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// POST routes

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, world!")
}

// GET routes

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	props := struct {
		NavbarProps NavbarProps
		Title       string
	}{
		NavbarProps: NavbarProps{
			Links: []Link{
				{RouteName: "Test", URL: "/test"},
			},
		},
		Title: "Go-HTMX Home",
	}

	tmpl, err := template.ParseFiles("./public/index.html", "./templates/navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, props)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	props := struct {
		NavbarProps NavbarProps
		Title       string
	}{
		NavbarProps: NavbarProps{
			Links: []Link{
				{RouteName: "Home", URL: "/"},
			},
		},
		Title: "Go-HTMX Test",
	}

	tmpl, err := template.ParseFiles("./public/test.html", "./templates/navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, props)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
