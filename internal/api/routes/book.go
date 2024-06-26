package routes

import (
	"github.com/afurgapil/library-management-system/internal/api/handlers"
	"github.com/afurgapil/library-management-system/internal/api/middleware"
	"github.com/afurgapil/library-management-system/pkg/book"

	"github.com/gofiber/fiber/v2"
)

func BookRouter(app fiber.Router, service book.Service) {
	app.Post("/add",middleware.DevelopmentEmployeeTokenMiddleware, handlers.AddBook(service))
	app.Delete("/delete/:id",middleware.DevelopmentStudentTokenMiddleware,handlers.DeleteBook(service))
	app.Get("/get/:id",handlers.GetBook(service))
	app.Get("/gets",handlers.GetBooks(service))
	app.Get("/gets-by-ID",handlers.GetBooksByIDHandler(service))
}