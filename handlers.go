package main

import (
	"context"
	"html/template"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoItem struct {
	Id         int
	ItemString string
}

// Common queries

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

// DELETE routes

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

	QueryAllTodos(w, r, dbpool)
}

// POST routes

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

	QueryAllTodos(w, r, dbpool)
}

func FetchTodoHandler(w http.ResponseWriter, r *http.Request, dbpool *pgxpool.Pool) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	QueryAllTodos(w, r, dbpool)
}

// GET routes

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	props := struct {
		NavbarProps NavbarProps
		Title       string
		Todos       []TodoItem
	}{
		NavbarProps: NavbarProps{
			Links: []Link{
				{RouteName: "Test", URL: "/test"},
			},
		},
		Title: "Go-HTMX Home",
		Todos: []TodoItem{},
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
