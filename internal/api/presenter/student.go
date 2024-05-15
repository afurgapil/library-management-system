package presenter

import (
	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type StudentResponse struct {
    StudentID       string 	`json:"student_id"`
    StudentMail     string 	`json:"student_mail"`
    StudentPassword string 	`json:"student_password"`
    Debit           int 	`json:"debit"`
    BookLimit 		int 	`json:"book_limit"`
    IsBanned       	bool 	`json:"isBanned"`
}

func StudentSuccessResponse(data *entities.Student) *fiber.Map {
	student := StudentResponse{
		StudentID:            data.StudentID,
		StudentMail:         data.StudentMail,
		Debit:         data.Debit,
		BookLimit:data.BookLimit,
		IsBanned:     data.IsBanned,
	}
	return &fiber.Map{
		"status":true,
		"data":student,
		"error":nil,
	}
}

func StudentsSuccessResponse(data *[]StudentResponse) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func StudentErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}