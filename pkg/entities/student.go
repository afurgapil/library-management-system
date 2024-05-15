package entities

type Student struct {
    StudentID       string 	`json:"student_id"`
    StudentMail     string 	`json:"student_mail"`
    StudentPassword string 	`json:"student_password"`
    Debit           int 	`json:"debit"`
    BookLimit 		int 	`json:"book_limit"`
    IsBanned       	bool 	`json:"isBanned"`
}
