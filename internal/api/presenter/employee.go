package presenter

import (
	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type EmployeeResponse struct {
    EmployeeID         	string `json:"employee_id"`
    EmployeeMail       	string `json:"employee_mail"`
    EmployeeUsername    string `json:"employee_username"`
    EmployeePassword	string `json:"employee_password"`
    EmployeePhoneNumber	string `json:"employee_phone_number"`
    Position			string `json:"position"`
}

func EmployeeSuccessResponse(data *entities.Employee) *fiber.Map {
	employee:=EmployeeResponse{
		EmployeeID:            	data.EmployeeID,
		EmployeeMail:         	data.EmployeeMail,
		EmployeeUsername:       data.EmployeeUsername,
		EmployeePassword:       data.EmployeePassword,
		EmployeePhoneNumber:	data.EmployeePhoneNumber,
		Position:     			data.Position,
	}
	return &fiber.Map{
		"status":true,
		"data":employee,
		"error":nil,
	}
}

func EmployeesSuccessResponse(data *[]EmployeeResponse) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func EmployeeErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}


