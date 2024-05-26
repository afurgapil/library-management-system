package routes

import (
	"github.com/afurgapil/library-management-system/internal/api/handlers"
	"github.com/afurgapil/library-management-system/internal/api/middleware"
	"github.com/afurgapil/library-management-system/pkg/student"

	"github.com/gofiber/fiber/v2"
)

func StudentRouter(app fiber.Router, service student.Service) {
	app.Post("/add",middleware.DevelopmentStudentTokenMiddleware, handlers.AddStudent(service))
	app.Post("/signin", handlers.StudentSignIn(service))
	app.Post("/password-reset-request", handlers.RequestPasswordResetHandler(service))
	app.Post("/reset-password/:token", handlers.ResetPasswordHandler(service))
	app.Post("/borrow-book",middleware.DevelopmentStudentTokenMiddleware,handlers.BookBorrowHandler(service))
	app.Post("/deliver-book/:borrowID/:bookID/:studentID",middleware.DevelopmentStudentTokenMiddleware,handlers.DeliverBookHandler(service))
	app.Post("/extend/:borrowID",middleware.DevelopmentStudentTokenMiddleware,handlers.ExtendDateHandler(service))
	app.Get("/get-borrowed-books/:studentID",middleware.DevelopmentStudentTokenMiddleware,handlers.GetBorrowedBooksHandler(service))
}