package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	envFile := "../../.env.development"

	if err := godotenv.Load(envFile); err != nil {
		log.Println("No .env.dev file found")
	}
}

func Connect() (*pgx.Conn, error) {

	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid port number: %v", err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database.")
	return conn, nil
}
