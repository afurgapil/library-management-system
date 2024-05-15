package entities

type Employee struct {
    EmployeeID         	string `json:"employee_id"`
    EmployeeMail       	string `json:"employee_mail"`
    EmployeeUsername    string `json:"employee_username"`
    EmployeePassword	string `json:"employee_password"`
    EmployeePhoneNumber	string `json:"employee_phone_number"`
    Position			string `json:"position"`
}
