package testutils

import (
	"context"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/jackc/pgx/v4"
)

var ExampleBook = &entities.Book{
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
}

var ExampleStudent = &entities.Student{
	StudentID:       "student_id",
	StudentMail:     "student_mail",
	StudentPassword: "student_password",
	Debit:           20,
	BookLimit:       20,
	IsBanned:        false,
}

var ExampleBorrowedBook = &entities.BorrowedBook{
	BorrowID:     "borrow_id",
	StudentID:    "student_id",
	BookID:       "book_id",
	BorrowDate:   "borrow_date",
	DeliveryDate: "delivery_date",
	IsExtended:   false,
}

func SetupTestDataBook(db *pgx.Conn, book *entities.Book) error {
	_, err := db.Exec(context.Background(), `INSERT INTO book (book_id, title, author, genre, publication_date, publisher, isbn, page_count, shelf_number, language, donor) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, book.BookID, book.Title, book.Author, book.Genre, book.PublicationDate, book.Publisher, book.ISBN, book.PageCount, book.ShelfNumber, book.Language, book.Donor)
	return err
}

func CleanupTestDataBook(db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), `DELETE FROM book`)
	return err
}

func SetupTestDataStudent(db *pgx.Conn, student *entities.Student) error {
	_, err := db.Exec(context.Background(), `INSERT INTO student (student_id, student_mail, student_password, debit, book_limit, is_banned) 
              VALUES ($1, $2, $3, $4, $5, $6)`, student.StudentID, student.StudentMail, student.StudentPassword, student.Debit, student.BookLimit, student.IsBanned)
	return err
}

func CleanupTestDataStudent(db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), `DELETE FROM student`)
	return err
}

func SetupTestDataBookBorrow(db *pgx.Conn, bookBorrow *entities.BorrowedBook) error {
	_, err := db.Exec(context.Background(), `INSERT INTO book_borrow (borrow_id, student_id, book_id, borrow_date, delivery_date, is_extended) 
              VALUES ($1, $2, $3, $4, $5, $6)`, bookBorrow.BorrowID, bookBorrow.StudentID, bookBorrow.BookID, bookBorrow.BorrowDate, bookBorrow.DeliveryDate, bookBorrow.IsExtended)
	return err
}

func CleanupTestDataBookBorrow(db *pgx.Conn) error {
	_, err := db.Exec(context.Background(), `DELETE FROM book_borrow`)
	return err
}
