package student

import (
	"context"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/utils"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
	AddStudent(student *entities.Student) (*entities.Student, error)
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

    query := `INSERT INTO student (student_id, student_mail, student_password, debit, book_limit, isBanned) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING student_id`
    err = r.DB.QueryRow(context.Background(), query,
        studentID, student.StudentMail, student.StudentPassword, student.Debit, student.BookLimit, student.IsBanned).Scan(&student.StudentID)
    if err != nil {
        return nil, err
    }

    return student, nil
}
