package get_handlers

import (
	"html/template"
	"htmx-go/helper_functions"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	props := struct {
		NavbarProps helper_functions.NavbarProps
		Title       string
		Todos       []helper_functions.TodoItem
	}{
		NavbarProps: helper_functions.NavbarProps{
			Links: []helper_functions.Link{
				{RouteName: "Test", URL: "/test"},
			},
		},
		Title: "Go-HTMX Home",
		Todos: []helper_functions.TodoItem{},
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
