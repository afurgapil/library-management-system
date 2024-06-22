package testutils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DbConnection *pgx.Conn

func init() {
	envFile := "../../.env.test"

	if err := godotenv.Load(envFile); err != nil {
		log.Println("No .env.test file found")
	}
}
func InitializeDatabase() (*pgx.Conn, error) {
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

	DbConnection = conn
	log.Println("Successfully connected to the database.")
	return DbConnection, nil
}

func CloseDatabase() {
	if DbConnection != nil {
		DbConnection.Close(context.Background())
		DbConnection = nil
	}
}

func TestMain(m *testing.M) {
	if _, err := InitializeDatabase(); err != nil {
		log.Fatalf("could not initialize database: %v", err)
	}
	defer CloseDatabase()

	code := m.Run()
	os.Exit(code)
}
