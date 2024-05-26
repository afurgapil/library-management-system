package entities

type BorrowedBook struct {
	BorrowID		string `json:"borrow_id"`
    StudentID       string `json:"student_id"`
	BookID 			string `json:"book_id"`
	BorrowDate 		string `json:"borrow_date"`
	DeliveryDate 	string `json:"delivery_date"`
	IsExtended 		bool `json:"is_extended"`
}