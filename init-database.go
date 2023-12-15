package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofor-little/env"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDb() (*pgxpool.Pool, error) {
	err := env.Load("./postgres/.env")
	if err != nil {
		log.Fatal(err)
	}

	dbURL := "postgresql://postgres:" + os.Getenv("POSTGRES_PASSWORD") + "@localhost:5432/postgres"

	dbpool, err := pgxpool.New(context.Background(), dbURL)

	return dbpool, err
}

func CreateTableIfNotExists(dbpool *pgxpool.Pool) {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS todo (
			id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
			todoItem VARCHAR(255)
		)
	`
	_, err := dbpool.Exec(context.Background(), createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully/already exists!")
}
