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
    IsBanned       	bool 	`json:"is_banned"`
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

func StudentInvalidRequestPayload() *fiber.Map  {
	return &fiber.Map{
		"status":false,
		"data": "",
		"error":"invalid request payload",
	}
}

func StudentInvalidRequestParams() *fiber.Map  {
	return &fiber.Map{
		"status":false,
		"data": "",
		"error":"missing required parameters",
	}
}

func StudentInternalServerError(err error) *fiber.Map  {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error": "Internal Server Error:"+ err.Error(),
	}
}

func StudentOKResponse(msg string) *fiber.Map  {
	return &fiber.Map{
		"status": true,
		"data":   msg,
		"error":  nil,
	}
}

func StudentGetBorrowedBooksOK(data *[]entities.BorrowedBook) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}
