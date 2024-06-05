package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "2412"
	dbname   = "librarymanagementsystem"
)

func Connect() (*pgx.Conn, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database.")
	return conn, nil
}
