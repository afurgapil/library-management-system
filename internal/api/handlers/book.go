package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/afurgapil/library-management-system/internal/api/presenter"
	"github.com/afurgapil/library-management-system/pkg/book"
	"github.com/afurgapil/library-management-system/pkg/entities"
)

// AddBook godoc
// @Summary Add a new book
// @Description add by json book
// @Tags book
// @Accept  json
// @Produce  json
// @Param   book  body      entities.Book   true  "Add Book"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /book/add [post]
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

// DeleteEmployee godoc
// @Summary Delete a book
// @Description Delete a book by ID
// @Tags book
// @Accept  json
// @Produce  json
// @Param   id   path      string   true  "Book ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /book/delete/{id} [delete]
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
