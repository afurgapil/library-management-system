package employee

import (
	"errors"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/utils"
)

type Service interface {
	InsertEmployee(employee *entities.Employee) (*entities.Employee, error)
	SignIn(email, password string) (string, *entities.Employee, error)
	DeleteEmployee(employeeID string) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) InsertEmployee(employee *entities.Employee) (*entities.Employee, error) {
	return s.repo.CreateEmployee(employee)
}

func (s *service) SignIn(email, password string) (string, *entities.Employee, error) {
	employee, err := s.repo.AuthenticateUser(email, password)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(employee.EmployeeID)
	if err != nil {
		return "", nil, err
	}

	return token, employee, nil
}

func (s *service) DeleteEmployee(employeeID string) error {
	err:=s.repo.DeleteEmployee(employeeID)
	if err != nil {
		return err
	}
	return nil
}