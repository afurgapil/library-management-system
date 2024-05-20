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
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
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
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /student/signin [post]
func StudentSignIn(service student.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&request); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.StudentErrorResponse(errors.New(
		 		"invalid request")))
		}

		token, student, err := service.SignIn(request.Email, request.Password)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(presenter.StudentErrorResponse(errors.New(
		 		"invalid email or password")))
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
// @Param request body map[string]string true "Email"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /password-reset-request [post]
func RequestPasswordResetHandler(service student.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		type request struct {
			Email string `json:"email"`
		}

		var req request
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status": false,
				"error":  "invalid request payload",
			})
		}

		err := service.RequestPasswordReset(req.Email)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"status": false,
				"error":  err.Error(),
			})
		}

		c.Status(http.StatusOK)
		return c.JSON(fiber.Map{
			"status":  true,
			"message": "password reset email sent",
		})
	}
}

// ResetPasswordHandler godoc
// @Summary Reset password
// @Description Reset the student's password using a token
// @Tags student
// @Accept json
// @Produce json
// @Param request body map[string]string true "Token and New Password"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /reset-password [post]
func ResetPasswordHandler(service student.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
        token := c.Params("token") 

		type request struct {
			NewPassword string `json:"new_password"`
		}

		var req request
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status": false,
				"error":  "invalid request payload",
			})
		}	

		err := service.ResetPassword(token, req.NewPassword)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"status": false,
				"error":  err.Error(),
			})
		}

		c.Status(http.StatusOK)
		return c.JSON(fiber.Map{
			"status":  true,
			"message": "password reset successful",
		})
	}
}