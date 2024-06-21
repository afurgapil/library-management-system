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
	if !utils.CheckMailValid(employee.EmployeeMail) {
		return nil, errors.New("invalid email address")
	}
	if len(employee.EmployeeUsername) < 8 {
		return nil, errors.New("username should be minimum 8 characters")
	}
	if len(employee.EmployeeUsername) > 64 {
		return nil, errors.New("username should be maximum 64 characters")
	}
	if len(employee.EmployeePassword) < 8 {
		return nil, errors.New("password should be minimum 8 characters")
	}
	if len(employee.EmployeePassword) > 64 {
		return nil, errors.New("password should be maximum 64 characters")
	}

	newEmployee, err := s.repo.CreateEmployee(employee)
	if err != nil {
		return nil, err
	}

	return newEmployee, nil
}

func (s *service) SignIn(email, password string) (string, *entities.Employee, error) {
	if email == "" {
		return "", nil, errors.New("email cannot be null")
	}
	if password == "" {
		return "", nil, errors.New("password cannot be null")
	}

	employee, err := s.repo.AuthenticateUser(email, password)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateEmployeeJWT(employee.EmployeeID)
	if err != nil {
		return "", nil, err
	}

	return token, employee, nil
}

func (s *service) DeleteEmployee(employeeID string) error {
	if employeeID == "" {
		return errors.New("employee ID cannot be empty")
	}
	err := s.repo.DeleteEmployee(employeeID)
	if err != nil {
		return err
	}
	return nil
}
