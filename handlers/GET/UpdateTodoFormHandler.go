package get

import (
	"context"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UpdateTodoFormHandler(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
	idStr := mux.Vars(r)["id"]

	transaction, err := dbpool.Begin(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer transaction.Rollback(context.Background())

	rows, err := transaction.Query(context.Background(), "SELECT id, todoItem FROM todo WHERE id=($1)", idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todoItems []TodoItem

	for rows.Next() {
		var todoItemId int
		var todoItemString string
		if err := rows.Scan(&todoItemId, &todoItemString); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todoItems = append(todoItems, TodoItem{Id: todoItemId, ItemString: todoItemString})
	}

	err = transaction.Commit(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	props := struct {
		NavbarProps NavbarProps
		Title       string
		Todo        TodoItem
	}{
		NavbarProps: NavbarProps{
			Links: []Link{
				{RouteName: "Test", URL: "/test"},
				{RouteName: "Home", URL: "/"},
			},
		},
		Title: "Go-HTMX Update Todo",
		Todo:  todoItems[0],
	}

	tmpl, err := template.ParseFiles("./public/updatetodo.html", "./templates/navbar.html")
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
