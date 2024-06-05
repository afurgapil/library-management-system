package utils

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
)

func setupTestDataCheckStudentMailExist(db *pgx.Conn, email string) error {
	_, err := db.Exec(context.Background(), `INSERT INTO student (student_id, student_mail, student_password, debit, book_limit,is_banned ) 
              VALUES ($1, $2, $3, $4, $5, $6)`, "0", email, "a", 42, 42, false)
	return err
}

func cleanupTestDataCheckStudentMailExist(db *pgx.Conn, studentMail string) error {
	_, err := db.Exec(context.Background(), `DELETE FROM student WHERE student_mail = $1`, studentMail)
	return err
}

func TestCheckStudentMailExist(t *testing.T) {
	tests := []struct {
		name    string
		mail    string
		want    bool
		wantErr bool
	}{
		{
			name:    "Mail Exist",
			mail:    "asd@gmail.com",
			want:    true,
			wantErr: true,
		}, {
			name:    "Mail Not Exist",
			mail:    "a@gmail.com",
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Mail Exist" {
				if err := setupTestDataCheckStudentMailExist(dbConnection, tt.mail); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}

			defer cleanupTestDataCheckStudentMailExist(dbConnection, tt.mail)

			got, err := CheckStudentMailExist(dbConnection, tt.mail)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckMailExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckMailExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
