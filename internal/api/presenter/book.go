package presenter

import (
	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type BookResponse struct {
	ID             string `json:"string"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Genre          string `json:"genre"`
	PublicationDate string `json:"publication_date"`
	Publisher      string `json:"publisher"`
	ISBN           string `json:"isbn"`
	PageCount      int    `json:"page_count"`
	ShelfNumber    string `json:"shelf_number"`
	Language       string `json:"language"`
	Donor          string `json:"donor"`
}

func BookSuccessResponse(data *entities.Book) *fiber.Map {
	book:=BookResponse{
		ID:            data.BookID,
		Title:         data.Title,
		Author:        data.Author,
		Genre:         data.Genre,
		PublicationDate:data.PublicationDate,
		Publisher:     data.Publisher,
		ISBN:          data.ISBN,
		PageCount:     data.PageCount,
		ShelfNumber:   data.ShelfNumber,
		Language:      data.Language,
		Donor:         data.Donor,	
	}
	return &fiber.Map{
		"status":true,
		"data":book,
		"error":nil,
	}
}

func BooksSuccessResponse(data *[]BookResponse) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}