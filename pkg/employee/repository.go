package employee

import (
	"context"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/afurgapil/library-management-system/pkg/utils"
	"github.com/jackc/pgx/v4"
)

type Repository interface {
    CreateEmployee(employee *entities.Employee) (*entities.Employee, error)
	AuthenticateUser(email, password string) (*entities.Employee, error)
    DeleteEmployee(employeeID string) error

}


type repository struct {
	DB *pgx.Conn
}

func NewRepo(db *pgx.Conn) Repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateEmployee(employee *entities.Employee) (*entities.Employee, error) {
    var employeeID string
    var isUnique bool
    var err error

  
    for {
        employeeID = utils.GenerateUniqueID()
        isUnique, err = utils.CheckIdValue(r.DB, "employee", "employee_id", employeeID)
        if err != nil {
            return nil, err
        }
        if isUnique {
            break
        }
    }

    query := `INSERT INTO employee (employee_id, employee_mail, employee_username, employee_phone_number, position, employee_password) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING employee_id`
    err = r.DB.QueryRow(context.Background(), query,
        employeeID, employee.EmployeeMail, employee.EmployeeUsername, employee.EmployeePhoneNumber, employee.Position, employee.EmployeePassword).Scan(&employee.EmployeeID)

    if err != nil {
        return nil, err
    }

    return employee, nil
}


func (r *repository) AuthenticateUser(email, password string) (*entities.Employee, error) {
	var employee entities.Employee
	query := `SELECT employee_id, employee_mail, employee_username, employee_phone_number, position, employee_password FROM employee WHERE employee_mail = $1 AND employee_password = $2`

	err := r.DB.QueryRow(context.Background(), query, email, password).Scan(
		&employee.EmployeeID,
		&employee.EmployeeMail,
		&employee.EmployeeUsername,
		&employee.EmployeePhoneNumber,
		&employee.Position,
		&employee.EmployeePassword,
	)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *repository) DeleteEmployee(employeeID string) error {
    query:="DELETE FROM employee WHERE employee_id = $1"
    _, err := r.DB.Exec(context.Background(),query,employeeID)
    if err != nil {
        return err
    }
    return nil
}