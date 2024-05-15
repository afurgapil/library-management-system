package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func CheckIdValue(db *pgx.Conn, tableName, columnName, idValue string) (bool, error) {
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = $1", tableName, columnName)

    var count int
    err := db.QueryRow(context.Background(), query, idValue).Scan(&count)
    if err != nil {
        return false, err 
    }

    return count == 0, nil
}
