package routes

import (
	"github.com/afurgapil/library-management-system/internal/api/handlers"
	"github.com/afurgapil/library-management-system/internal/api/middleware"
	"github.com/afurgapil/library-management-system/pkg/employee"

	"github.com/gofiber/fiber/v2"
)

func EmployeeRouter(app fiber.Router, service employee.Service) {
	app.Post("/add", handlers.AddEmployee(service))
	app.Post("/signin", handlers.SignIn(service))
	app.Delete("/delete/:id",middleware.DevelopmentTokenMiddleware ,handlers.DeleteEmployee(service))

}