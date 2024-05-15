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
