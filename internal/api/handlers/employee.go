package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/afurgapil/library-management-system/internal/api/presenter"
	"github.com/afurgapil/library-management-system/pkg/employee"
	"github.com/afurgapil/library-management-system/pkg/entities"
)



func AddEmployee(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Employee
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		if requestBody.EmployeeMail == "" || requestBody.EmployeeUsername == "" {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(errors.New(
				"employee email and username are required")))
		}
		result, err := service.InsertEmployee(&requestBody)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		c.Status(fiber.StatusCreated)
		return c.JSON(presenter.EmployeeSuccessResponse(result))
	}
}


func SignIn(service employee.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&request); err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenter.EmployeeErrorResponse(errors.New(
		 		"invalid request")))
		}

		token, employee, err := service.SignIn(request.Email, request.Password)
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(presenter.EmployeeErrorResponse(errors.New(
		 		"invalid email or password")))
		}

		return c.JSON(fiber.Map{
			"status":   true,
			"token":    token,
			"employee": presenter.EmployeeSuccessResponse(employee),
		})
	}
}

func DeleteEmployee(service employee.Service) fiber.Handler  {
	return func(c *fiber.Ctx) error {
		employeeID := c.Params("id")
		if employeeID =="" {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{"error":"Employee ID is required"})
		}
		err := service.DeleteEmployee(employeeID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.EmployeeErrorResponse(err))
		}
		return c.SendStatus(http.StatusNoContent)
	}
}