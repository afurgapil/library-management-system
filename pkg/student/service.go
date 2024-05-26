package student

import (
	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/utils"
)

type Service interface {
    InsertStudent(student *entities.Student) (*entities.Student, error)
    SignIn(email, password string) (string, *entities.Student, error)
    RequestPasswordReset(email string) error
    ResetPassword(token, newPassword string) error
    BookBorrow(bookID, studentID string) (string, error) 
    DeliverBook(borrowID,bookID,studentID string) (string,error)
	ExtendDate(borrowID string) (string,error)
	GetBorrowedBooks(studentID string) ([]*entities.BorrowedBook, error)

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
        return "", nil, err
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

func (s *service) BookBorrow(bookID, studentID string) (string, error) {
    borrowID, err := s.repo.BorrowBook(bookID, studentID)
    if err != nil {
        return "", err
    }

    return borrowID, nil
}

func (s *service) DeliverBook(borrowId,bookId,studentID string) (string,error)  {
    msg,err:=s.repo.DeliverBook(borrowId,bookId,studentID)
    if err != nil {
        return "", err
    }
    
    return msg, nil
}

func (s *service) ExtendDate(borrowID string) (string,error)  {
    msg,err := s.repo.ExtendDate(borrowID)
    if err !=nil {
        return "",err
    }
    return msg,nil
}

func (s *service) GetBorrowedBooks(studentID string) ([]*entities.BorrowedBook, error) {
    borrowedBooks, err := s.repo.GetBorrowedBooks(studentID)
    if err != nil {
        return nil, err
    }
    
    borrowedBooksPtrs := make([]*entities.BorrowedBook, len(borrowedBooks))
    for i, book := range borrowedBooks {
        borrowedBooksPtrs[i] = &book
    }
    
    return borrowedBooksPtrs, nil
}
