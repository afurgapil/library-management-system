package testutils

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

var DbConnection *pgx.Conn

func InitializeDatabase() (*pgx.Conn, error) {
	var err error
	DbConnection, err = pgx.Connect(context.Background(), "postgres://postgres:test1234@localhost:5432/librarymanagementsystem_test")
	if err != nil {
		return nil, err
	}
	return DbConnection, nil
}

func CloseDatabase() {
	if DbConnection != nil {
		DbConnection.Close(context.Background())
	}
}

func TestMain(m *testing.M) {
	var err error
	DbConnection, err = InitializeDatabase()
	if err != nil {
		panic(err)
	}
	defer CloseDatabase()

	code := m.Run()
	os.Exit(code)
}
