package book

import (
	"context"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/utils"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
    CreateBook(book *entities.Book) (*entities.Book, error)
    DeleteBook(bookID string) error
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
		bookID=utils.GenerateUniqueID()
		isUnique,err=utils.CheckIdValue(r.DB,"book","book_id",bookID)
		if err != nil {
            return nil, err  
        }
        if isUnique {
            break 
        }
	}

	query := `INSERT INTO book (book_id, title, author, genre, publication_date, publisher, isbn, page_count, shelf_number, language, donor) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING book_id`
	err = r.DB.QueryRow(context.Background(), query,
		bookID, book.Title, book.Author, book.Genre, book.PublicationDate, book.Publisher, book.ISBN, book.PageCount, book.ShelfNumber, book.Language, book.Donor).Scan(&book.BookID)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *repository) DeleteBook(bookID string) error {
    query := "DELETE FROM book WHERE book_id = $1"
    _, err := r.DB.Exec(context.Background(), query, bookID)
    if err != nil {
        return err
    }
    return nil
}
