package book

import (
	"errors"

	"github.com/afurgapil/library-management-system/pkg/entities"
)

type Service interface {
	InsertBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(bookID string) error
	GetBook(bookID string) (*entities.Book, error)
	GetBooks() ([]*entities.Book, error)
	GetBooksByID(bookIDList []string) ([]*entities.Book, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) InsertBook(book *entities.Book) (*entities.Book, error) {
	return s.repo.CreateBook(book)
}

func (s *service) DeleteBook(bookID string) error {
	if bookID == "" {
		return errors.New("book ID cannot be empty")
	}
	return s.repo.DeleteBook(bookID)
}

func (s *service) GetBook(bookID string) (*entities.Book, error) {
	if bookID == "" {
		return nil, errors.New("book ID cannot be empty")
	}
	return s.repo.GetBook(bookID)

}

func (s *service) GetBooks() ([]*entities.Book, error) {
	return s.repo.GetBooks()
}
func (s *service) GetBooksByID(bookIdList []string) ([]*entities.Book, error) {
	return s.repo.GetBooksByID(bookIdList)
}
