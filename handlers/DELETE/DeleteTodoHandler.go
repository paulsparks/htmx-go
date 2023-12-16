package delete

import (
	"context"
	"htmx-go/reusedresponses"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := mux.Vars(r)["id"]

	transaction, err := dbpool.Begin(context.Background())
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer transaction.Rollback(context.Background())

	_, err = transaction.Exec(context.Background(), "DELETE FROM todo WHERE id=($1)", idStr)
	if err != nil {
		http.Error(w, "Error while deleting from database", http.StatusInternalServerError)
	}

	err = transaction.Commit(context.Background())
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
	}

	reusedresponses.RespondWithTodos(w, r, dbpool)
}
