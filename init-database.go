package main

import (
	"context"
	"log"
	"os"

	"github.com/gofor-little/env"
	"github.com/jackc/pgx/v5"
)

func ConnectToDb() (*pgx.Conn, error) {
	err := env.Load("./postgres/.env")
	if err != nil {
		log.Fatal(err)
	}

	dbURL := "postgresql://postgres:" + os.Getenv("POSTGRES_PASSWORD") + "@localhost:5432/postgres"

	conn, err := pgx.Connect(context.Background(), dbURL)

	return conn, err
}
