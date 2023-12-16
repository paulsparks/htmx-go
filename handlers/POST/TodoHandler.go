package post

import (
	"context"
	"htmx-go/reusedqueries"
	"net/http"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TodoHandler(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	form := url.Values{}

	err := r.ParseForm()
	if err != nil {
		form.Set("todoItem", "Error")
	}

	form = r.Form

	transaction, err := dbpool.Begin(context.Background())
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer transaction.Rollback(context.Background())

	_, err = transaction.Exec(context.Background(), "INSERT INTO todo (todoItem) VALUES ($1)", form.Get("todoItem"))
	if err != nil {
		http.Error(w, "Error while inserting into database", http.StatusInternalServerError)
	}

	err = transaction.Commit(context.Background())
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
	}

	reusedqueries.QueryAllTodos(w, r, dbpool)
}
