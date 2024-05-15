package student

import "github.com/afurgapil/library-management-system/pkg/entities"

type Service interface {
	InsertStudent(student *entities.Student) (*entities.Student, error)
}

type service struct {
	repo Repository 
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) InsertStudent(student *entities.Student) (*entities.Student, error) {
	return s.repo.AddStudent(student)
}
