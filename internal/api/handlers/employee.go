package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/afurgapil/library-management-system/internal/api/presenter"
	"github.com/afurgapil/library-management-system/pkg/employee"
	"github.com/afurgapil/library-management-system/pkg/entities"
)

// AddEmployee godoc
// @Summary Add a new employee
// @Description add by json employee
// @Tags employee
// @Accept  json
// @Produce  json
// @Param   employee  body      entities.Employee   true  "Add Employee"
// @Success 201 {object} presenter.EmployeeSuccessResponseStruct
// @Failure 400 {object} presenter.EmployeeErrorResponseStruct
// @Router /employee/add [post]
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

// SignIn godoc
// @Summary Sign in an employee
// @Description Sign in an employee with email and password
// @Tags employee
// @Accept  json
// @Produce  json
// @Param   signin  body      object{email=string,password=string}   true  "Sign In"
// @Success 200 {object} presenter.EmployeeSuccessResponseStruct
// @Failure 400 {object} presenter.EmployeeErrorResponseStruct
// @Failure 401 {object} presenter.EmployeeErrorResponseStruct
// @Router /employee/signin [post]
func EmployeeSignIn(service employee.Service) fiber.Handler {
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

// DeleteEmployee godoc
// @Summary Delete an employee
// @Description Delete an employee by ID
// @Tags employee
// @Accept  json
// @Produce  json
// @Param   id   path      string   true  "Employee ID"
// @Success 204 "No Content"
// @Failure 400 {object} presenter.EmployeeErrorResponseStruct
// @Failure 500 {object} presenter.EmployeeErrorResponseStruct
// @Router /employee/delete/{id} [delete]
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