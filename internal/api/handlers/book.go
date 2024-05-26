package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/afurgapil/library-management-system/internal/api/presenter"
	"github.com/afurgapil/library-management-system/pkg/book"
	"github.com/afurgapil/library-management-system/pkg/entities"
)

// @Summary Add a new book
// @Description add by json book
// @Tags book
// @Accept  json
// @Produce  json
// @Param   book  body      entities.Book   true  "Add Book"
// @Success 201 {object} map[string]interface{} "Book creation success response"
// @Failure 400 {object} map[string]interface{} "Error response"
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

// @Summary Delete a book
// @Description Delete a book by ID
// @Tags book
// @Accept  json
// @Produce  json
// @Param   id   path      string   true  "Book ID"
// @Success 204
// @Failure 400 {object} map[string]interface{} "Error response"
// @Failure 500 {object} map[string]interface{} "Internal server error"
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

// @Summary Get a book
// @Description Get a book by ID
// @Tags book
// @Accept  json
// @Produce  json
// @Param   id   path      string   true  "Book ID"
// @Success 200 {object} map[string]interface{} "Book retrieval success response"
// @Failure 400 {object} map[string]interface{} "Error response"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /book/get/{id} [get]
func GetBook(service book.Service) fiber.Handler  {
	return func(c *fiber.Ctx) error {
		bookID := c.Params("id")
		if bookID == "" {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{"error":"Book ID is required"})
		}
		result,err := service.GetBook(bookID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		c.Status(http.StatusOK)
		return c.JSON(presenter.BookSuccessResponse(result))
	}
}


// GetBooks godoc
// @Summary Get a listed books
// @Description Get a listed books with an order
// @Tags book
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /book/gets [get]
func GetBooks(service book.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := service.GetBooks()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
		}
		c.Status(http.StatusOK)
		return c.JSON(presenter.BooksSuccessResponse(result))
	}
}


// @Summary Get books by ID
// @Description Retrieve a list of books by their IDs
// @Tags books
// @Accept json
// @Produce json
// @Param request body object true "Book ID List"
// @Success 200 {object} map[string]interface{}{"status":bool,"data":[]entities.Book}
// @Failure 400 {object} map[string]interface{}{"status":bool,"error":string}
// @Failure 500 {object} map[string]interface{}{"status":bool,"error":string}
// @Router /book/ [post]
func GetBooksByIDHandler(service book.Service) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var request struct {
            BookIDs []string `json:"book_ids"`
        }

        if err := c.BodyParser(&request); err != nil {
            c.Status(http.StatusBadRequest)
			return c.JSON(presenter.BookErrorResponse(err))
        }

        books, err := service.GetBooksByID(request.BookIDs)
        if err != nil {
           	c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.BookErrorResponse(err))
        }

        c.Status(http.StatusOK)
		return c.JSON(presenter.BooksSuccessResponse(books))
    }
}
