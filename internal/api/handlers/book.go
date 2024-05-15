package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/afurgapil/library-management-system/internal/api/presenter"
	"github.com/afurgapil/library-management-system/pkg/book"
	"github.com/afurgapil/library-management-system/pkg/entities"
)



func AddBook(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Book
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		if requestBody.Author == "" || requestBody.Title == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(errors.New(
				"please specify title and author")))
		}
		result, err := service.InsertBook(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		c.Status(http.StatusCreated)
		return c.JSON(presenter.BookSuccessResponse(result))
	}
}

func DeleteBook(service book.Service) fiber.Handler {
    return func(c *fiber.Ctx) error {
        bookID := c.Params("id") 
		if bookID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{"error": "Book ID is required"})
		}
        err := service.DeleteBook(bookID)
        if err != nil {
            c.Status(http.StatusInternalServerError)
            return c.JSON(presenter.BookErrorResponse(err))
        }
        return c.SendStatus(http.StatusNoContent)
    }
}
