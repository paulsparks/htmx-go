package put

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := mux.Vars(r)["id"]

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

	_, err = transaction.Exec(context.Background(), "UPDATE todo SET todoItem=($1) WHERE id=($2)", form.Get("todoItem"), idStr)
	if err != nil {
		http.Error(w, "Error while updating database", http.StatusInternalServerError)
	}

	err = transaction.Commit(context.Background())
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
	}

	w.Header().Set("HX-Redirect", "/")

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprint(w, `{"message": "Success"}`)
}
