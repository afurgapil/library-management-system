package student

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
	"github.com/afurgapil/library-management-system/pkg/utils"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

// Student - DB Connection
var dbConnection *pgx.Conn

func TestMain(m *testing.M) {
	var err error
	dbConnection, err = testutils.InitializeDatabase()
	if err != nil {
		panic(err)
	}
	defer testutils.CloseDatabase()

	code := m.Run()
	os.Exit(code)
}
func Test_repository_AddStudent(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		student *entities.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Student
		wantErr bool
	}{
		{

			name: "Added Student Succesfully",
			args: args{
				testutils.ExampleStudent,
			},
			want:    testutils.ExampleStudent,
			wantErr: false,
		},
		{
			name: "Student Mail Exists",
			fields: fields{
				DB: nil,
			},
			args: args{
				student: testutils.ExampleStudent,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Student Mail Exists" {
				if err := testutils.SetupTestDataStudent(dbConnection, tt.args.student); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}
			defer testutils.CleanupTestDataStudent(dbConnection)

			r := &repository{
				DB: dbConnection,
			}
			got, err := r.AddStudent(tt.args.student)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.AddStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.AddStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_AuthenticateStudent(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Student
		wantErr bool
	}{
		{
			name: "Authenticated Student Successfully",
			args: args{
				email:    "student@mail.com",
				password: "student_password",
			},
			want: &entities.Student{
				StudentID:       "student_id",
				StudentMail:     "student@mail.com",
				StudentPassword: "$2a$10$k.ImSVXMyQec0y4tI1l9UOWvVzWYQwOEzHWplWvEyUTnwQ0NX32mu", // Hashed password
				Debit:           20,
				BookLimit:       20,
				IsBanned:        false,
			},
			wantErr: false,
		},
		{
			name: "Invalid Credentials",
			args: args{
				email:    "test@example.com",
				password: "wrongpassword",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Student Banned",
			args: args{
				email:    "test@example.com",
				password: "testpassword",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, err := utils.EncryptPassword(tt.args.password)
			if err != nil {
				t.Fatalf("Failed to encrypt: %v", err)
			}
			switch tt.name {
			case "Authenticated Student Successfully":
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student@mail.com",
					StudentPassword: hashedPassword,
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			case "Student Banned":
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student@mail.com",
					StudentPassword: hashedPassword,
					Debit:           20,
					BookLimit:       20,
					IsBanned:        true,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}

			defer testutils.CleanupTestDataStudent(dbConnection)

			r := &repository{
				DB: dbConnection,
			}
			got, err := r.AuthenticateStudent(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.AuthenticateStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				tt.want.StudentPassword = got.StudentPassword // Set hashed password dynamically
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.AuthenticateStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetStudentByEmail(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Student
		wantErr bool
	}{
		{
			name: "Student Fetched Successfully",
			args: args{
				email: "test@example.com",
			},
			want: &entities.Student{
				StudentID:   "student_id",
				StudentMail: "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "Student Not Found",
			args: args{
				email: "nonexistent@example.com",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Student Fetched Successfully" {
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:   "student_id",
					StudentMail: "test@example.com",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}
			defer testutils.CleanupTestDataStudent(dbConnection)

			r := &repository{
				DB: dbConnection,
			}
			got, err := r.GetStudentByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetStudentByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetStudentByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_UpdateStudentPassword(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		studentID      string
		hashedPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Update Student Password Successfully",
			args: args{
				studentID:      "student_id",
				hashedPassword: "new_student_password",
			},
			wantErr: false,
		},
		{
			name: "Update Student Password Error",
			args: args{
				studentID:      "non-existent-id",
				hashedPassword: "hashed-password",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Update Student Password Successfully":
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}

			defer testutils.CleanupTestDataStudent(dbConnection)

			r := &repository{
				DB: dbConnection,
			}
			err := r.UpdateStudentPassword(tt.args.studentID, tt.args.hashedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdateStudentPassword() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.name == "UpdateStudentPasswordSuccess" {
				var student entities.Student
				query := `SELECT student_password FROM student WHERE student_id = $1`
				err = dbConnection.QueryRow(context.Background(), query, tt.args.studentID).Scan(&student.StudentPassword)
				if err != nil {
					t.Errorf("Error getting student password after update: %v", err)
				}
				if student.StudentPassword != tt.args.hashedPassword {
					t.Errorf("Password was not updated correctly. Got: %v, Want: %v", student.StudentPassword, tt.args.hashedPassword)
				}
			}
		})
	}
}

func Test_repository_BorrowBook(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		bookID    string
		studentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Book Borrowed Successfully",
			args: args{
				bookID:    "book_id",
				studentID: "student_id",
			},
			want:    "borrow_id",
			wantErr: false,
		},
		{
			name: "Debit Error",
			args: args{
				bookID:    "book_id",
				studentID: "student_id",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "No Limit Error",
			args: args{
				bookID:    "book_id",
				studentID: "student_id",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Borrowed Already Error",
			args: args{
				bookID:    "book_id",
				studentID: "student_id",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Book Borrowed Successfully":
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           0,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBook(dbConnection, &entities.Book{
					BookID:          "book_id",
					Title:           "book_title",
					Author:          "book_Author",
					Genre:           "book_genre",
					PublicationDate: "99.99.9999",
					Publisher:       "book_publisher",
					ISBN:            "book_isbn",
					PageCount:       20,
					ShelfNumber:     "book_shelf",
					Language:        "book_language",
					Donor:           "book_donor",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}

			case "Debit Error":
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBook(dbConnection, &entities.Book{
					BookID:          "book_id",
					Title:           "book_title",
					Author:          "book_Author",
					Genre:           "book_genre",
					PublicationDate: "99.99.9999",
					Publisher:       "book_publisher",
					ISBN:            "book_isbn",
					PageCount:       20,
					ShelfNumber:     "book_shelf",
					Language:        "book_language",
					Donor:           "book_donor",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}

			case "No Limit Error":
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBook(dbConnection, &entities.Book{
					BookID: "book_id",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}

			case "Borrowed Already Error":
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBook(dbConnection, &entities.Book{
					BookID: "book_id",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBookBorrow(dbConnection, &entities.BorrowedBook{
					BorrowID:  "borrow_id",
					BookID:    "book_id",
					StudentID: "student_id",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}

			defer testutils.CleanupTestDataBookBorrow(dbConnection)
			defer testutils.CleanupTestDataStudent(dbConnection)
			defer testutils.CleanupTestDataBook(dbConnection)

			r := &repository{
				DB: dbConnection,
			}
			got, err := r.BorrowBook(tt.args.bookID, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.BorrowBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "Book Borrowed Successfully" && got == "" {
				t.Errorf("repository.BorrowBook(),BorrowBookSuccess = %v, want %v", got, tt.want)

			}
			if got != tt.want && tt.name != "Book Borrowed Successfully" {
				t.Errorf("repository.BorrowBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_DeliverBook(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		borrowID  string
		bookID    string
		studentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Book Delivered Successfully",
			args: args{
				borrowID:  "borrow_id",
				bookID:    "book_id",
				studentID: "student_id",
			},
			want:    "Book returned successfully",
			wantErr: false,
		},
		{
			name: "Borrow Record Not Found",
			args: args{
				borrowID:  "borrow_id",
				bookID:    "book_id",
				studentID: "student_id",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Book Delivered overdue",
			args: args{
				borrowID:  "borrow_id",
				bookID:    "book_id",
				studentID: "student_id",
			},
			want:    "Book returned successfully",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Book Delivered Successfully":
				currentTime := time.Now()
				borrowDate := currentTime.AddDate(0, -1, 0).Format("2006-01-02")
				deliveryDate := currentTime.AddDate(0, 1, 0).Format("2006-01-02")
				if err := testutils.SetupTestDataBook(dbConnection, &entities.Book{
					BookID:          "book_id",
					Title:           "book_title",
					Author:          "book_Author",
					Genre:           "book_genre",
					PublicationDate: "99.99.9999",
					Publisher:       "book_publisher",
					ISBN:            "book_isbn",
					PageCount:       20,
					ShelfNumber:     "book_shelf",
					Language:        "book_language",
					Donor:           "book_donor",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBookBorrow(dbConnection, &entities.BorrowedBook{
					BorrowID:     "borrow_id",
					StudentID:    "student_id",
					BookID:       "book_id",
					BorrowDate:   borrowDate,
					DeliveryDate: deliveryDate,
					IsExtended:   false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			case "Book Delivered overdue":
				currentTime := time.Now()
				borrowDate := currentTime.AddDate(0, -2, 0).Format("2006-01-02")
				deliveryDate := currentTime.AddDate(0, -1, 0).Format("2006-01-02")
				if err := testutils.SetupTestDataBook(dbConnection, &entities.Book{
					BookID:          "book_id",
					Title:           "book_title",
					Author:          "book_Author",
					Genre:           "book_genre",
					PublicationDate: "99.99.9999",
					Publisher:       "book_publisher",
					ISBN:            "book_isbn",
					PageCount:       20,
					ShelfNumber:     "book_shelf",
					Language:        "book_language",
					Donor:           "book_donor",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBookBorrow(dbConnection, &entities.BorrowedBook{
					BorrowID:     "borrow_id",
					BookID:       "book_id",
					StudentID:    "student_id",
					BorrowDate:   borrowDate,
					DeliveryDate: deliveryDate,
					IsExtended:   false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}

			}

			defer testutils.CleanupTestDataBookBorrow(dbConnection)
			defer testutils.CleanupTestDataStudent(dbConnection)
			defer testutils.CleanupTestDataBook(dbConnection)

			r := &repository{
				DB: dbConnection,
			}
			got, err := r.DeliverBook(tt.args.borrowID, tt.args.bookID, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.DeliverBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repository.DeliverBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_ExtendDate(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		borrowID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Date Extended Successfully",
			args: args{
				borrowID: "borrow_id",
			},
			want:    "date extended successfully",
			wantErr: false,
		},
		{
			name: "Borrow Record Not Found",
			args: args{
				borrowID: "borrow_id",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Date Extended Already",
			args: args{
				borrowID: "borrow_id",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Date Extended Successfully":
				currentTime := time.Now()
				borrowDate := currentTime.AddDate(0, -1, 0).Format("2006-01-02")
				deliveryDate := currentTime.AddDate(0, 1, 0).Format("2006-01-02")
				if err := testutils.SetupTestDataBook(dbConnection, &entities.Book{
					BookID:          "book_id",
					Title:           "book_title",
					Author:          "book_Author",
					Genre:           "book_genre",
					PublicationDate: "99.99.9999",
					Publisher:       "book_publisher",
					ISBN:            "book_isbn",
					PageCount:       20,
					ShelfNumber:     "book_shelf",
					Language:        "book_language",
					Donor:           "book_donor",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBookBorrow(dbConnection, &entities.BorrowedBook{
					BorrowID:     "borrow_id",
					StudentID:    "student_id",
					BookID:       "book_id",
					BorrowDate:   borrowDate,
					DeliveryDate: deliveryDate,
					IsExtended:   false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			case "Date Extended Already":
				currentTime := time.Now()
				borrowDate := currentTime.AddDate(0, -1, 0).Format("2006-01-02")
				deliveryDate := currentTime.AddDate(0, 1, 0).Format("2006-01-02")
				if err := testutils.SetupTestDataBook(dbConnection, &entities.Book{
					BookID:          "book_id",
					Title:           "book_title",
					Author:          "book_Author",
					Genre:           "book_genre",
					PublicationDate: "99.99.9999",
					Publisher:       "book_publisher",
					ISBN:            "book_isbn",
					PageCount:       20,
					ShelfNumber:     "book_shelf",
					Language:        "book_language",
					Donor:           "book_donor",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBookBorrow(dbConnection, &entities.BorrowedBook{
					BorrowID:     "borrow_id",
					StudentID:    "student_id",
					BookID:       "book_id",
					BorrowDate:   borrowDate,
					DeliveryDate: deliveryDate,
					IsExtended:   true,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}

			defer testutils.CleanupTestDataBookBorrow(dbConnection)
			defer testutils.CleanupTestDataStudent(dbConnection)
			defer testutils.CleanupTestDataBook(dbConnection)

			r := &repository{
				DB: dbConnection,
			}
			got, err := r.ExtendDate(tt.args.borrowID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.ExtendDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repository.ExtendDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetBorrowedBooks(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		studentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.BorrowedBook
		wantErr bool
	}{
		{
			name: "Books Fetched Successfully",
			args: args{
				studentID: "student_id",
			},
			want: []entities.BorrowedBook{
				{
					BorrowID:     "borrow_id_1",
					BookID:       "book_id_1",
					StudentID:    "student_id",
					BorrowDate:   "2024-05-22",
					DeliveryDate: "2024-07-22",
					IsExtended:   false,
				},
				{
					BorrowID:     "borrow_id_2",
					BookID:       "book_id_2",
					StudentID:    "student_id",
					BorrowDate:   "2024-05-22",
					DeliveryDate: "2024-07-22",
					IsExtended:   false,
				},
			},
			wantErr: false,
		},
		{
			name: "Books Not Found",
			args: args{
				studentID: "student_id_null",
			},
			want:    []entities.BorrowedBook{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Books Fetched Successfully":
				if err := testutils.SetupTestDataBook(dbConnection, &entities.Book{
					BookID:          "book_id_1",
					Title:           "book_title",
					Author:          "book_Author",
					Genre:           "book_genre",
					PublicationDate: "99.99.9999",
					Publisher:       "book_publisher",
					ISBN:            "book_isbn_1",
					PageCount:       20,
					ShelfNumber:     "book_shelf",
					Language:        "book_language",
					Donor:           "book_donor",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBook(dbConnection, &entities.Book{
					BookID:          "book_id_2",
					Title:           "book_title",
					Author:          "book_Author",
					Genre:           "book_genre",
					PublicationDate: "99.99.9999",
					Publisher:       "book_publisher",
					ISBN:            "book_isbn_2",
					PageCount:       20,
					ShelfNumber:     "book_shelf",
					Language:        "book_language",
					Donor:           "book_donor",
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBookBorrow(dbConnection, &entities.BorrowedBook{
					BorrowID:     "borrow_id_1",
					StudentID:    "student_id",
					BookID:       "book_id_1",
					BorrowDate:   "2024-05-22",
					DeliveryDate: "2024-07-22",
					IsExtended:   false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				if err := testutils.SetupTestDataBookBorrow(dbConnection, &entities.BorrowedBook{
					BorrowID:     "borrow_id_2",
					StudentID:    "student_id",
					BookID:       "book_id_2",
					BorrowDate:   "2024-05-22",
					DeliveryDate: "2024-07-22",
					IsExtended:   false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			case "Books Not Found":
				if err := testutils.SetupTestDataStudent(dbConnection, &entities.Student{
					StudentID:       "student_id_null",
					StudentMail:     "student_mail",
					StudentPassword: "student_password",
					Debit:           20,
					BookLimit:       20,
					IsBanned:        false,
				}); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}

			defer testutils.CleanupTestDataBookBorrow(dbConnection)
			defer testutils.CleanupTestDataStudent(dbConnection)
			defer testutils.CleanupTestDataBook(dbConnection)

			r := &repository{
				DB: dbConnection,
			}
			got, err := r.GetBorrowedBooks(tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetBorrowedBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "Books Not Found" {
				tt.want = got
			}
			if len(got) != 0 && tt.name == "Books Not Found" {
				t.Errorf("repository.GetBorrowedBooks() = %v, want %v", got, tt.want)

			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetBorrowedBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}
