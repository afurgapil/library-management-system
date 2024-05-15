package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/afurgapil/library-management-system/internal/api/presenter"
	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/student"
)



func AddStudent(service student.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Student
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.StudentErrorResponse(err))
		}
		if requestBody.StudentMail == "" || requestBody.StudentPassword == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.StudentErrorResponse(errors.New(
				"please specify mail and password")))
		}
		result, err := service.InsertStudent(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.StudentErrorResponse(err))
		}
		return c.JSON(presenter.StudentSuccessResponse(result))
	}
}