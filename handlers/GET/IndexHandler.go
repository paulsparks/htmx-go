package get

import (
	"html/template"
	"net/http"
)

type Link struct {
	RouteName string
	URL       string
}

type NavbarProps struct {
	Links []Link
}

type TodoItem struct {
	Id         int
	ItemString string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	props := struct {
		NavbarProps NavbarProps
		Title       string
		Todos       []TodoItem
	}{
		NavbarProps: NavbarProps{
			Links: []Link{
				{RouteName: "Test", URL: "/test"},
			},
		},
		Title: "Go-HTMX Home",
		Todos: []TodoItem{},
	}

	tmpl, err := template.ParseFiles("./public/index.html", "./templates/navbar.html", "./templates/todo-list.html")
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
