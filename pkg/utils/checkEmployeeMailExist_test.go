package utils

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

func setupTestDataCheckEmployeeMailExist(db *pgx.Conn, email string) error {
	_, err := db.Exec(context.Background(), `INSERT INTO employee (employee_id, employee_mail, employee_username, employee_phone_number, position, employee_password) 
              VALUES ($1, $2, $3, $4, $5, $6)`, "0", email, "a", "b", "c", "d")
	return err
}

func cleanupTestDataCheckEmployeeMailExist(db *pgx.Conn, employeeMail string) error {
	_, err := db.Exec(context.Background(), `DELETE FROM employee WHERE employee_mail = $1`, employeeMail)
	return err
}

func TestCheckEmployeeMailExist(t *testing.T) {
	tests := []struct {
		name           string
		mail           string
		want           bool
		wantErr        bool
		wantErrMessage string
	}{
		{
			name:           "Mail Exist",
			mail:           "asd@gmail.com",
			want:           true,
			wantErr:        true,
			wantErrMessage: "this mail has been used already",
		},
		{
			name:           "Mail Not Exist",
			mail:           "a@gmail.com",
			want:           false,
			wantErr:        false,
			wantErrMessage: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Mail Exist" {
				if err := setupTestDataCheckEmployeeMailExist(dbConnection, tt.mail); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}

			defer cleanupTestDataCheckEmployeeMailExist(dbConnection, tt.mail)

			got, err := CheckEmployeeMailExist(dbConnection, tt.mail)
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
