package reusedresponses

import (
	"context"
	gethandler "htmx-go/handlers/GET"
	"net/http"
	"text/template"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RespondWithTodos(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
	transaction, err := dbpool.Begin(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer transaction.Rollback(context.Background())

	rows, err := transaction.Query(context.Background(), "SELECT id, todoItem FROM todo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todoItems []gethandler.TodoItem

	for rows.Next() {
		var todoItemId int
		var todoItemString string
		if err := rows.Scan(&todoItemId, &todoItemString); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todoItems = append(todoItems, gethandler.TodoItem{Id: todoItemId, ItemString: todoItemString})
	}

	err = transaction.Commit(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("./templates/todo-list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	props := struct {
		Todos []gethandler.TodoItem
	}{
		Todos: todoItems,
	}

	w.WriteHeader(http.StatusOK)

	err = tmpl.ExecuteTemplate(w, "todo-list", props)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
