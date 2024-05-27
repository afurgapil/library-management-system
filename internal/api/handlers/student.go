package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/afurgapil/library-management-system/internal/api/presenter"
	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/student"
)

// AddStudent godoc
// @Summary Add a new student
// @Description add by json student
// @Tags student
// @Accept  json
// @Produce  json
// @Param   student  body      entities.Student   true  "Add Student"
// @Success 201 {object} presenter.StudentSuccessResponseStruct
// @Failure 400 {object} presenter.StudentErrorResponseStruct
// @Failure 500 {object} presenter.StudentErrorResponseStruct
// @Router /student/add [post]
func AddStudent(service student.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Student
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.StudentErrorResponse(err))
		}
		if requestBody.StudentMail == "" || requestBody.StudentPassword == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.StudentErrorResponse(errors.New(
				"student email and password are required")))
		}
		result, err := service.InsertStudent(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.StudentErrorResponse(err))
		}
		c.Status(fiber.StatusCreated)
		return c.JSON(presenter.StudentSuccessResponse(result))
	}
}

// SignIn godoc
// @Summary Sign in a student
// @Description Sign in a student with email and password
// @Tags student
// @Accept  json
// @Produce  json
// @Param   signin  body      object{email=string,password=string}   true  "Sign In"
// @Success 201 {object} presenter.StudentSuccessResponseStruct
// @Failure 400 {object} presenter.StudentErrorResponseStruct
// @Failure 401 {object} presenter.StudentErrorResponseStruct
// @Failure 500 {object} presenter.StudentErrorResponseStruct
// @Router /student/signin [post]
func StudentSignIn(service student.Service) fiber.Handler {
    return func(c *fiber.Ctx) error {
        if c == nil {
            return errors.New("context is nil")
        }
        
        var request struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }

        if err := c.BodyParser(&request); err != nil {
            c.Status(fiber.StatusBadRequest)
            return c.JSON(presenter.StudentErrorResponse(errors.New("invalid request")))
        }

        token, student, err := service.SignIn(request.Email, request.Password)
        if err != nil {
            switch err.Error() {
                case "student banned":
                    c.Status(fiber.StatusForbidden)
                    return c.JSON(presenter.StudentErrorResponse(errors.New(err.Error())))
                case "invalid email or password":
                    c.Status(fiber.StatusUnauthorized)
                    return c.JSON(presenter.StudentErrorResponse(errors.New(err.Error())))
                default:
                    c.Status(fiber.StatusInternalServerError)
                    return c.JSON(presenter.StudentErrorResponse(errors.New("err")))
            }
        }

        return c.JSON(fiber.Map{
            "status":   true,
            "token":    token,
            "student": presenter.StudentSuccessResponse(student),
        })
    }
}


// RequestPasswordResetHandler godoc
// @Summary Request password reset
// @Description Send a password reset email to the student
// @Tags student
// @Accept json
// @Produce json
// @Param resetRequest body object{email=string} true "Reset Request"
// @Success 201 {object} presenter.StudentOKResponseStruct
// @Failure 400 {object} presenter.StudentErrorResponseStruct
// @Failure 500 {object} presenter.StudentErrorResponseStruct
// @Router /student/password-reset-request [post]
func RequestPasswordResetHandler(service student.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			Email string `json:"email"`
		}

		var req request
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.StudentInvalidRequestPayload())
		}

		err := service.RequestPasswordReset(req.Email)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.StudentInternalServerError(err))

		}

		c.Status(http.StatusOK)
		return c.JSON(presenter.StudentOKResponse("password reset email sent"))

	}
}

// ResetPasswordHandler godoc
// @Summary Reset password
// @Description Reset the student's password using a token
// @Tags student
// @Accept json
// @Produce json
// @Param token path string true "Token"
// @Param request body object{new_email=string} true "New Password"
// @Success 200 {object} presenter.StudentOKResponseStruct
// @Failure 400 {object} presenter.StudentErrorResponseStruct
// @Failure 500 {object} presenter.StudentErrorResponseStruct
// @Router /student/reset-password/{token} [post]
func ResetPasswordHandler(service student.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
        token := c.Params("token") 

		type request struct {
			NewPassword string `json:"new_password"`
		}

		var req request
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.StudentInvalidRequestPayload())
		}	

		err := service.ResetPassword(token, req.NewPassword)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.StudentInternalServerError(err))
		}
		c.Status(http.StatusOK)
		return c.JSON(presenter.StudentOKResponse("password reset successful"))
	}
}

// BookBorrowHandler godoc
// @Summary Borrow a book
// @Description Allows a student to borrow a book
// @Tags student
// @Accept  json
// @Produce  json
// @Param   request  body presenter.StudentBorrowBookRequestStruct true "Borrow Book Request"
// @Success 200 {object} presenter.StudentOKResponseStruct
// @Failure 400 {object} presenter.StudentErrorResponseStruct
// @Failure 500 {object} presenter.StudentErrorResponseStruct
// @Router /student/borrow-book/ [post]
func BookBorrowHandler(service student.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type BorrowRequest struct {
			StudentID string `json:"student_id"`
			BookID    string `json:"book_id"`
		}

		var request BorrowRequest
		if err := c.BodyParser(&request); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.StudentInvalidRequestPayload())
		}

		borrowID, err := service.BookBorrow(request.BookID, request.StudentID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.StudentInternalServerError(err))
		}
		c.Status(http.StatusOK)
		return c.JSON(presenter.StudentOKResponse("book borrowed succesfully. borrow_id :"+borrowID))
	}
}


// DeliverBookHandler godoc
// @Summary Deliver a book
// @Description Allows a student to deliver a book
// @Tags student
// @Accept  json
// @Produce  json
// @Param   borrow_id  path string true "Borrow ID"
// @Param   book_id     path string true "Book ID"
// @Param   student_id  path string true "Student ID"
// @Success 200 {object} presenter.StudentOKResponseStruct
// @Failure 400 {object} presenter.StudentErrorResponseStruct
// @Failure 500 {object} presenter.StudentErrorResponseStruct
// @Router /student/deliver-book/{borrow_id}/{book_id}/{student_id} [post]
func DeliverBookHandler(service student.Service) fiber.Handler {
    return func(c *fiber.Ctx) error {
        borrowID := c.Params("borrowID")
        bookID := c.Params("bookID")
        studentID := c.Params("studentID")

        if borrowID == "" || bookID == "" || studentID == "" {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.StudentInvalidRequestParams())

        }

        message, err := service.DeliverBook(borrowID, bookID, studentID)
        if err != nil {
			c.Status(fiber.StatusInternalServerError)
            return c.JSON(presenter.StudentInternalServerError(err))
        }

        c.Status(fiber.StatusOK)
        return c.JSON(presenter.StudentOKResponse(message))
    }
}

// @Summary Extend Borrow Date
// @Description Extend the borrow date for a given borrow ID.
// @Tags student
// @Accept json
// @Produce json
// @Param borrowID path string true "Borrow ID"
// @Success 200 {object} presenter.StudentOKResponseStruct
// @Failure 400 {object} presenter.StudentErrorResponseStruct
// @Failure 500 {object} presenter.StudentErrorResponseStruct
// @Router /student/extend/{borrowID} [post]
func ExtendDateHandler(service student.Service) fiber.Handler {
    return func(c *fiber.Ctx) error {
        borrowID := c.Params("borrowID")

        if borrowID == "" {
			c.Status(fiber.StatusBadRequest)
         	return c.JSON(presenter.StudentInvalidRequestParams())
        }

        message, err := service.ExtendDate(borrowID)
        if err != nil {
			switch err.Error() {
			case  "date extended already":
				c.Status(fiber.StatusNoContent)
				return c.JSON(presenter.StudentErrorResponse(err))
			case  "borrow record does not exist":
				c.Status(fiber.StatusNotFound)
				return c.JSON(presenter.StudentErrorResponse(err))
			default:
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(presenter.StudentInternalServerError(err))
			}
        }
		c.Status(fiber.StatusOK)
        return c.JSON(presenter.StudentOKResponse(message))

    }
}


// @Summary Get Borrowed Books
// @Description Get a list of borrowed books for a given student ID
// @Tags student
// @Accept json
// @Produce json
// @Param studentID path string true "Student ID"
// @Success 200 {object} presenter.StudentBorrowedBookStruct
// @Failure 400 {object} presenter.StudentErrorResponseStruct
// @Failure 500 {object} presenter.StudentErrorResponseStruct
// @Router /students/get-borrowed-books/{studentID} [get]
func GetBorrowedBooksHandler(service student.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		studentID := c.Params("studentID")

		if studentID == "" {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.StudentInvalidRequestParams())
		}

		borrowedBooks, err := service.GetBorrowedBooks(studentID)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenter.StudentInternalServerError(err))
		}

		var books []entities.BorrowedBook
		for _, book := range borrowedBooks {
			books = append(books, *book)
		}

		c.Status(fiber.StatusOK)
		return c.JSON(presenter.StudentGetBorrowedBooksOK(&books))
	}
}