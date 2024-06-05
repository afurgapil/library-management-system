package student

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/utils"
	"github.com/jackc/pgx/v4"
)

// TODO add mail notifications && cron jobs
type Repository interface {
	AddStudent(student *entities.Student) (*entities.Student, error)
	AuthenticateStudent(email, password string) (*entities.Student, error)
	GetStudentByEmail(email string) (*entities.Student, error)
	UpdateStudentPassword(studentID, hashedPassword string) error
	BorrowBook(bookID, studentID string) (string, error)
	DeliverBook(borrowID, bookID, studentID string) (string, error)
	ExtendDate(borrowID string) (string, error)
	GetBorrowedBooks(studentID string) ([]entities.BorrowedBook, error)
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
	mailExist, mailErr := utils.CheckStudentMailExist(r.DB, student.StudentMail)
	if mailErr != nil || mailExist {
		return nil, errors.New("this mail has been used already")
	}
	for {
		studentID = utils.GenerateUniqueID()
		isUnique, err = utils.CheckIdValue(r.DB, "student", "student_id", studentID)
		if err != nil {
			return nil, err
		}
		if !isUnique {
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

	err = utils.DecryptPassword(student.StudentPassword, password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	studentStatus, statusErr := utils.CheckStudentBanStatus(r.DB, email)
	if statusErr != nil || studentStatus {
		return nil, errors.New("student banned")
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

func (r *repository) BorrowBook(bookID, studentID string) (string, error) {
	var borrowID string
	currentTime := time.Now()
	borrowDate := currentTime.Format("2006-01-02")
	deliveryDate := currentTime.AddDate(0, 1, 0).Format("2006-01-02")
	var isUnique bool
	var err error

	studentDebit, err1 := utils.CheckDebit(r.DB, studentID)
	studentBookLimit, err2 := utils.CheckLimit(r.DB, studentID)
	bookStatus, err3 := utils.CheckBookStatus(r.DB, bookID)
	if err1 != nil || err2 != nil || err3 != nil {
		return "", err
	}

	if !studentDebit {
		return "", fmt.Errorf("student should pay dabit: %w", err)
	}
	if !studentBookLimit {
		return "", fmt.Errorf("student have no book limit: %w", err)
	}
	if !bookStatus {
		return "", fmt.Errorf("book borrowed already: %w", err)
	}

	for {
		borrowID = utils.GenerateUniqueID()
		isUnique, err = utils.CheckIdValue(r.DB, "book_borrow", "borrow_id", borrowID)
		if err != nil {
			return "", err
		}
		if !isUnique {
			break
		}
	}

	query := `INSERT INTO book_borrow (borrow_id, student_id, book_id, borrow_date, delivery_date, is_extended) VALUES ($1, $2, $3, $4, $5, $6) RETURNING borrow_id`
	row := r.DB.QueryRow(context.Background(), query, borrowID, studentID, bookID, borrowDate, deliveryDate, false)

	var returnedBorrowID string
	err = row.Scan(&returnedBorrowID)
	if err != nil {
		return "", fmt.Errorf("error inserting borrow record: %w", err)
	}

	err = utils.UpdateLimit(r.DB, studentID, "borrow")
	if err != nil {
		return "", fmt.Errorf("error updating book limit: %w", err)
	}

	return returnedBorrowID, nil
}

func (r *repository) DeliverBook(borrowID, bookID, studentID string) (string, error) {
	isExist, err := utils.CheckIdValue(r.DB, "book_borrow", "borrow_id", borrowID)
	if err != nil {
		return "", err
	}
	if !isExist {
		return "", errors.New("borrow record does not exist")
	}

	daysRemaining, err := utils.CheckDate(r.DB, borrowID)
	if err != nil {
		return "", err
	}
	if daysRemaining < 0 {
		overdueDays := -daysRemaining
		penalty := overdueDays * 5
		err = utils.UpdateDebit(r.DB, studentID, penalty)
		if err != nil {
			return "", fmt.Errorf("error updating debit: %w", err)
		}
	}

	query := `DELETE FROM book_borrow WHERE borrow_id = $1 AND book_id = $2 AND student_id = $3`
	_, err = r.DB.Exec(context.Background(), query, borrowID, bookID, studentID)
	if err != nil {
		return "", fmt.Errorf("error deleting borrow record: %w", err)
	}

	err = utils.UpdateLimit(r.DB, studentID, "delivery")
	if err != nil {
		return "", fmt.Errorf("error updating book limit: %w", err)
	}
	return "Book returned successfully", nil
}

func (r *repository) ExtendDate(borrowID string) (string, error) {
	isExist, err := utils.CheckIdValue(r.DB, "book_borrow", "borrow_id", borrowID)
	if err != nil {
		return "", err
	}
	if !isExist {
		return "", errors.New("borrow record does not exist")
	}

	extendStatus, extendStatusErr := utils.CheckExtend(r.DB, borrowID)
	if extendStatusErr != nil {
		return "", extendStatusErr
	}
	if extendStatus {
		return "", errors.New("date extended already")
	}

	updateStatus, updateStatusErr := utils.UpdateExtend(r.DB, borrowID)
	if updateStatusErr != nil {
		return "", updateStatusErr
	}
	if !updateStatus {
		return "", errors.New("failed to extend date")
	}

	return "date extended successfully", nil
}

func (r *repository) GetBorrowedBooks(studentID string) ([]entities.BorrowedBook, error) {
	var borrowedBooks []entities.BorrowedBook

	query := `SELECT borrow_id, student_id, book_id, borrow_date, delivery_date, is_extended FROM book_borrow WHERE student_id = $1`
	rows, err := r.DB.Query(context.Background(), query, studentID)
	if err != nil {
		return nil, fmt.Errorf("error querying borrowed books: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var book entities.BorrowedBook
		err := rows.Scan(&book.BorrowID, &book.StudentID, &book.BookID, &book.BorrowDate, &book.DeliveryDate, &book.IsExtended)
		if err != nil {
			return nil, fmt.Errorf("error scanning borrowed book: %w", err)
		}
		borrowedBooks = append(borrowedBooks, book)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("error iterating through borrowed books: %w", rows.Err())
	}

	return borrowedBooks, nil
}
