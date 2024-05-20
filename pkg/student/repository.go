package student

import (
	"context"
	"errors"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/utils"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	AddStudent(student *entities.Student) (*entities.Student, error)
    AuthenticateStudent(email,password string) (*entities.Student, error)
    GetStudentByEmail(email string) (*entities.Student, error)
    UpdateStudentPassword(studentID, hashedPassword string) error
}

type repository struct {
	DB *pgx.Conn
}

func NewRepo(db *pgx.Conn) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) AddStudent(student *entities.Student) (*entities.Student, error) {
    var studentID string
    var isUnique bool
    var err error

    for {
        studentID = utils.GenerateUniqueID()
        isUnique, err = utils.CheckIdValue(r.DB, "student", "student_id", studentID)
        if err != nil {
            return nil, err  
        }
        if isUnique {
            break 
        }
    }
    hashedPassword, err := utils.EncryptPassword(student.StudentPassword)
        if err != nil {
            return nil, err
        }
    query := `INSERT INTO student (student_id, student_mail, student_password, debit, book_limit, is_banned) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING student_id`
    err = r.DB.QueryRow(context.Background(), query,
        studentID, student.StudentMail, hashedPassword, student.Debit, student.BookLimit, student.IsBanned).Scan(&student.StudentID)
    if err != nil {
        return nil, err
    }

    return student, nil
}

func (r *repository) AuthenticateStudent(email, password string) (*entities.Student, error) {
    var student entities.Student
    query := `SELECT student_id, student_mail, student_password, debit,book_limit, is_banned FROM student WHERE student_mail = $1`

    err :=r.DB.QueryRow(context.Background(),query,email).Scan(
        &student.StudentID,
        &student.StudentMail,
        &student.StudentPassword,
        &student.Debit,
        &student.BookLimit,
        &student.IsBanned,
    )
    if err != nil {
        return nil, err
    }

    err = utils.DecryptPassword(student.StudentPassword, password)
    if err != nil {
        return nil, errors.New("invalid email or password")
    }

    return &student, nil
}

func (r *repository) GetStudentByEmail(email string) (*entities.Student, error) {
    var student entities.Student
    query := `SELECT student_id, student_mail, student_password, debit, book_limit, is_banned FROM student WHERE student_mail = $1`

    err := r.DB.QueryRow(context.Background(), query, email).Scan(
        &student.StudentID,
        &student.StudentMail,
        &student.StudentPassword,
        &student.Debit,
        &student.BookLimit,
        &student.IsBanned,
    )
    if err != nil {
        return nil, err
    }

    return &student, nil
}

func (r *repository) UpdateStudentPassword(studentID, hashedPassword string) error {
    query := `UPDATE student SET student_password = $1 WHERE student_id = $2`

    _, err := r.DB.Exec(context.Background(), query, hashedPassword, studentID)
    if err != nil {
        return err
    }

    return nil
}
