package main

import (
	"log"

	_ "github.com/afurgapil/library-management-system/docs"
	"github.com/afurgapil/library-management-system/internal/api/routes"
	"github.com/afurgapil/library-management-system/internal/database"
	"github.com/afurgapil/library-management-system/pkg/book"
	"github.com/afurgapil/library-management-system/pkg/employee"
	"github.com/afurgapil/library-management-system/pkg/student"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture mongo book shop!"))
	})
	app.Get("/swagger/*", swagger.HandlerDefault)

	bookRepo := book.NewRepo(db)
	bookService := book.NewService(bookRepo)
	bookRoutes := app.Group("/api/book")

	studentRepo := student.NewRepo(db)
	studentService := student.NewService(studentRepo)
	studentRoutes := app.Group("/api/student")

	employeeRepo := employee.NewRepo(db)
	employeeService := employee.NewService(employeeRepo)
	employeeRoutes := app.Group("/api/employee")

	routes.BookRouter(bookRoutes, bookService)
	routes.StudentRouter(studentRoutes, studentService)
	routes.EmployeeRouter(employeeRoutes, employeeService)
	app.Listen(":3000")
}
