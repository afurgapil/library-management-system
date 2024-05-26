package utils

import (
	"context"

	"github.com/jackc/pgx/v4"
)

func CheckMailExist(db *pgx.Conn, mail string) (bool,error)  {
	var mailExists int

	query:= `SELECT COUNT(*) FROM student WHERE student_mail = $1`
	err:=db.QueryRow(context.Background(),query,mail).Scan(&mailExists)
	if err != nil {
		return false,err
	}
	return mailExists > 0,nil
}