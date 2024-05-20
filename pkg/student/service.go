package student

import (
	"errors"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/utils"
)

type Service interface {
    InsertStudent(student *entities.Student) (*entities.Student, error)
    SignIn(email, password string) (string, *entities.Student, error)
    RequestPasswordReset(email string) error
    ResetPassword(token, newPassword string) error
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

func (s *service) SignIn(email, password string) (string, *entities.Student, error) {
    student, err := s.repo.AuthenticateStudent(email, password)
    if err != nil {
        return "", nil, errors.New("invalid email or password")
    }
    token, err := utils.GenerateStudentJWT(student.StudentID)
    if err != nil {
        return "", nil, err
    }

    return token, student, nil
}

func (s *service) RequestPasswordReset(email string) error {
    student, err := s.repo.GetStudentByEmail(email)
    if err != nil {
        return err
    }

    token, err := utils.GeneratePasswordResetToken(student.StudentID)
    if err != nil {
        return err
    }

    err = utils.SendPasswordResetEmail(student.StudentMail, token)
    if err != nil {
        return err
    }

    return nil
}

func (s *service) ResetPassword(token, newPassword string) error {
    studentID, err := utils.ValidatePasswordResetToken(token)
    if err != nil {
        return err
    }

    hashedPassword, err := utils.EncryptPassword(newPassword)
    if err != nil {
        return err
    }

    err = s.repo.UpdateStudentPassword(studentID, hashedPassword)
    if err != nil {
        return err
    }

    return nil
}
