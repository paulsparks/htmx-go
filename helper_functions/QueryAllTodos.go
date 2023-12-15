package helper_functions

import (
	"context"
	"net/http"
	"text/template"

	"github.com/jackc/pgx/v5/pgxpool"
)

func QueryAllTodos(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
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

	var todoItems []TodoItem

	for rows.Next() {
		var todoItemId int
		var todoItemString string
		if err := rows.Scan(&todoItemId, &todoItemString); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todoItems = append(todoItems, TodoItem{todoItemId, todoItemString})
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
		Todos []TodoItem
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
