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

type BorrowedBook struct{
	BorrowID		string `json:"borrow_id"`
    StudentID       string `json:"student_id"`
	BookID 			string `json:"book_id"`
	BorrowDate 		string `json:"borrow_date"`
	DeliveryDate 	string `json:"delivery_date"`
	IsExtended 		bool `json:"is_extended"`
}

type StudentSuccessResponseStruct struct {
	Status bool `json:"status" example:"true"`
	Data StudentResponse `json:"data"`
	Error string `json:"error" example:""`
}

type StudentOKResponseStruct struct {
	Status bool `json:"status" example:"true"`
	Data string `json:"data"`
	Error string `json:"error" example:""`
}

type StudentBorrowedBookStruct struct {
	Status bool `json:"status" example:"true"`
	Data []BorrowedBook `json:"data"`
	Error string `json:"error" example:""`
}

type StudentBorrowBookRequestStruct struct {
	StudentID string `json:"student_id"`
	BookID string `json:"book_id"`
}

type StudentErrorResponseStruct struct {
	Status bool `json:"status" example:"false"`
	Data string `json:""`
	Error string `json:"error"`
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
