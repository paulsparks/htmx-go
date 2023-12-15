package get_handlers

import (
	"html/template"
	"htmx-go/helper_functions"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	props := struct {
		NavbarProps helper_functions.NavbarProps
		Title       string
	}{
		NavbarProps: helper_functions.NavbarProps{
			Links: []helper_functions.Link{
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
