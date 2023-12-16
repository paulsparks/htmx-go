package post

import (
	"htmx-go/reusedqueries"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func FetchTodoHandler(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	reusedqueries.QueryAllTodos(w, r, dbpool)
}
