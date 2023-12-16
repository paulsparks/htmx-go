package get

import (
	"html/template"
	"net/http"
)

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
