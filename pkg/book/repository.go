package book

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/utils"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	CreateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(bookID string) error
	GetBook(bookID string) (*entities.Book, error)
	GetBooks() ([]*entities.Book, error)
	GetBooksByID(bookIDList []string) ([]*entities.Book, error)
}

type repository struct {
	DB *pgx.Conn
}

func NewRepo(db *pgx.Conn) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateBook(book *entities.Book) (*entities.Book, error) {
	var bookID string
	var isUnique bool
	var err error

	for {
		bookID = utils.GenerateUniqueID()
		isUnique, err = utils.CheckIdValue(r.DB, "book", "book_id", bookID)
		if err != nil {
			return nil, err
		}
		if !isUnique {
			break
		}
	}

	query := `INSERT INTO book (book_id, title, author, genre, publication_date, publisher, isbn, page_count, shelf_number, language, donor) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING book_id`
	err = r.DB.QueryRow(context.Background(), query,
		book.BookID, book.Title, book.Author, book.Genre, book.PublicationDate, book.Publisher, book.ISBN, book.PageCount, book.ShelfNumber, book.Language, book.Donor).Scan(&book.BookID)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, errors.New("duplicated ID")
		}
		return nil, err
	}

	return book, nil
}

func (r *repository) DeleteBook(bookID string) error {
	query := "DELETE FROM book WHERE book_id = $1"

	result, err := r.DB.Exec(context.Background(), query, bookID)

	if err != nil {
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("book not found")
	}
	return nil
}

func (r *repository) GetBook(bookID string) (*entities.Book, error) {

	query := "SELECT * FROM book WHERE book_id = $1"

	row := r.DB.QueryRow(context.Background(), query, bookID)
	var book entities.Book
	err := row.Scan(
		&book.BookID,
		&book.Title,
		&book.Author,
		&book.Genre,
		&book.PublicationDate,
		&book.Publisher,
		&book.ISBN,
		&book.PageCount,
		&book.ShelfNumber,
		&book.Language,
		&book.Donor,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	return &book, nil
}

func (r *repository) GetBooks() ([]*entities.Book, error) {
	query := "SELECT * FROM book"
	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]*entities.Book, 0)
	for rows.Next() {
		var book entities.Book
		err := rows.Scan(
			&book.BookID,
			&book.Title,
			&book.Author,
			&book.Genre,
			&book.PublicationDate,
			&book.Publisher,
			&book.ISBN,
			&book.PageCount,
			&book.ShelfNumber,
			&book.Language,
			&book.Donor,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repository) GetBooksByID(bookIDList []string) ([]*entities.Book, error) {
	var books []*entities.Book
	for _, id := range bookIDList {
		if id == "" {
			return nil, errors.New("book ID list contains empty strings")
		}
	}

	query := `SELECT * FROM book WHERE book_id = ANY($1)`

	rows, err := r.DB.Query(context.Background(), query, bookIDList)
	if err != nil {
		return nil, fmt.Errorf("error querying books by ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book entities.Book
		err := rows.Scan(
			&book.BookID,
			&book.Title,
			&book.Author,
			&book.Genre,
			&book.PublicationDate,
			&book.Publisher,
			&book.ISBN,
			&book.PageCount,
			&book.ShelfNumber,
			&book.Language,
			&book.Donor)

		if err != nil {
			return nil, fmt.Errorf("error scanning book: %w", err)
		}
		books = append(books, &book)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterating through books: %w", rows.Err())
	}
	if len(books) == 0 {
		return nil, errors.New("books not found")

	}
	if len(bookIDList) != len(books) {
		return nil, errors.New("number of returned books does not match expected")
	}
	return books, nil
}
