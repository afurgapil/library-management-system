package testutils

import (
	"context"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/jackc/pgx/v4"
)

// utils
func SetupTestDataCreateBook(db *pgx.Conn, book *entities.Book) error {
	_, err := db.Exec(context.Background(), `INSERT INTO book (book_id, title, author, genre, publication_date, publisher, isbn, page_count, shelf_number, language, donor) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, book.BookID, book.Title, book.Author, book.Genre, book.PublicationDate, book.Publisher, book.ISBN, book.PageCount, book.ShelfNumber, book.Language, book.Donor)
	return err
}

func CleanupTestDataCreateBook(db *pgx.Conn, bookID string) error {
	_, err := db.Exec(context.Background(), `DELETE FROM book WHERE book_id = $1`, bookID)
	return err
}

func SetupTestDataDeleteBook(db *pgx.Conn, bookID string) error {
	_, err := db.Exec(context.Background(), `INSERT INTO book (book_id, title, author, genre, publication_date, publisher, isbn, page_count, shelf_number, language, donor) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, bookID, "book_title", "book_Author", "book_genre", "99-99-9999", "book_publisher", "book_isbn", 20, "book_shelf", "book_language", "book_donor")
	return err
}

func CleanupTestDataDeleteBook(db *pgx.Conn, bookID string) error {
	_, err := db.Exec(context.Background(), `DELETE FROM book WHERE book_id = $1`, bookID)
	return err
}

func SetupTestDataGetBook(db *pgx.Conn, bookID string) error {
	_, err := db.Exec(context.Background(), `INSERT INTO book (book_id, title, author, genre, publication_date, publisher, isbn, page_count, shelf_number, language, donor) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, bookID, "book_title", "book_Author", "book_genre", "99-99-9999", "book_publisher", "book_isbn", 20, "book_shelf", "book_language", "book_donor")
	return err
}

func CleanupTestDataGetBook(db *pgx.Conn, bookID string) error {
	_, err := db.Exec(context.Background(), `DELETE FROM book WHERE book_id = $1`, bookID)
	return err
}

func SetupTestDataGetBooks(db *pgx.Conn, books []*entities.Book) error {
	for _, book := range books {
		_, err := db.Exec(context.Background(), `INSERT INTO book (book_id, title, author, genre, publication_date, publisher, isbn, page_count, shelf_number, language, donor) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			book.BookID, book.Title, book.Author, book.Genre, book.PublicationDate, book.Publisher, book.ISBN, book.PageCount, book.ShelfNumber, book.Language, book.Donor)
		if err != nil {
			return err
		}
	}
	return nil
}

func CleanupTestDataGetBooks(db *pgx.Conn, books []*entities.Book) error {
	for _, book := range books {
		_, err := db.Exec(context.Background(), `DELETE FROM book WHERE book_id = $1`, book.BookID)
		if err != nil {
			return err
		}
	}
	return nil
}
