package book

import "github.com/afurgapil/library-management-system/pkg/entities"

type Service interface {
	InsertBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(bookID string) error
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
    err := s.repo.DeleteBook(bookID)
    if err != nil {
        return err
    }
    return nil
}