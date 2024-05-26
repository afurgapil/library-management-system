package utils

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func CheckBookStatus(db *pgx.Conn, bookID string) (bool,error)  {
	var bookCount int
	query := `SELECT COUNT(*) FROM book_borrow WHERE book_id = $1`	

	err:=db.QueryRow(context.Background(),query,bookID).Scan(&bookCount)
	if err != nil {
		return false,err
	}
	return bookCount <= 0,nil
}