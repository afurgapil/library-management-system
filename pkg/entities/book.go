package entities

type Book struct {
	BookID          string `json:"book_id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Genre           string `json:"genre"`
	PublicationDate string `json:"publication_date"`
	Publisher       string `json:"publisher"`
	ISBN            string `json:"isbn"`
	PageCount       int    `json:"page_count"`
	ShelfNumber     string `json:"shelf_number"`
	Language        string `json:"language"`
	Donor           string `json:"donor"`
}
