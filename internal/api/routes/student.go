package routes

import (
	"github.com/afurgapil/library-management-system/internal/api/handlers"
	"github.com/afurgapil/library-management-system/internal/api/middleware"
	"github.com/afurgapil/library-management-system/pkg/student"

	"github.com/gofiber/fiber/v2"
)

func StudentRouter(app fiber.Router, service student.Service) {
	app.Post("/add-student",middleware.DevelopmentTokenMiddleware, handlers.AddStudent(service))
}