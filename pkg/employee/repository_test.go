package employee

import (
	"os"
	"reflect"
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
	"github.com/jackc/pgx/v4"
)

// Employee - DB Connection
var dbConnection *pgx.Conn

func TestMain(m *testing.M) {
	var err error
	dbConnection, err = testutils.InitializeDatabase()
	if err != nil {
		panic(err)
	}
	defer testutils.CloseDatabase()

	code := m.Run()
	os.Exit(code)
}

func Test_repository_CreateEmployee(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		employee *entities.Employee
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           *entities.Employee
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "Create Employee Successfully",
			args: args{
				employee: &entities.Employee{
					EmployeeMail:        "test@example.com",
					EmployeeUsername:    "testuser",
					EmployeePassword:    "testpassword",
					EmployeePhoneNumber: "123456789",
					Position:            "tester",
				},
			},
			want: &entities.Employee{
				EmployeeMail:        "test@example.com",
				EmployeeUsername:    "testuser",
				EmployeePassword:    "testpassword",
				EmployeePhoneNumber: "123456789",
				Position:            "tester",
			},
			wantErr:        false,
			wantErrMessage: "",
		},
		{
			name: "Invalid Mail",
			args: args{
				employee: &entities.Employee{
					EmployeeMail:        "invalidmail",
					EmployeeUsername:    "testuser",
					EmployeePassword:    "testpassword",
					EmployeePhoneNumber: "123456789",
					Position:            "tester",
				},
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "invalid mail address",
		},
		{
			name: "Duplicate Mail",
			args: args{
				employee: &entities.Employee{
					EmployeeMail:        "duplicate@example.com",
					EmployeeUsername:    "testuser",
					EmployeePassword:    "testpassword",
					EmployeePhoneNumber: "123456789",
					Position:            "tester",
				},
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "this mail has been used already",
		},
		{
			name: "Short Username",
			args: args{
				employee: &entities.Employee{
					EmployeeMail:        "test@example.com",
					EmployeeUsername:    "user",
					EmployeePassword:    "testpassword",
					EmployeePhoneNumber: "123456789",
					Position:            "tester",
				},
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "username should be minimum 8 characters",
		},
		{
			name: "Long Username",
			args: args{
				employee: &entities.Employee{
					EmployeeMail:        "test@example.com",
					EmployeeUsername:    "usernamethatiswaytoolongandshouldnotbeacceptedbecauseitiswaymorethan64characterslong",
					EmployeePassword:    "testpassword",
					EmployeePhoneNumber: "123456789",
					Position:            "tester",
				},
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "username should be maximum 64 characters",
		},
		{
			name: "Short Password",
			args: args{
				employee: &entities.Employee{
					EmployeeMail:        "test@example.com",
					EmployeeUsername:    "testuser",
					EmployeePassword:    "short",
					EmployeePhoneNumber: "123456789",
					Position:            "tester",
				},
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "password should be minimum 8 characters",
		},
		{
			name: "Long Password",
			args: args{
				employee: &entities.Employee{
					EmployeeMail:        "test@example.com",
					EmployeeUsername:    "testuser",
					EmployeePassword:    "passwordthatiswaytoolongandshouldnotbeacceptedbecauseitiswaymorethan64characterslong",
					EmployeePhoneNumber: "123456789",
					Position:            "tester",
				},
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "password should be maximum 64 characters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Duplicate Mail" {
				if err := testutils.SetupTestDataCreateEmployee(dbConnection, tt.args.employee.EmployeeMail); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}
			defer testutils.CleanupTestDataCreateEmployee(dbConnection, tt.args.employee.EmployeeMail)

			r := &repository{
				DB: dbConnection,
			}
			got, err := r.CreateEmployee(tt.args.employee)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err == nil || err.Error() != tt.wantErrMessage {
					t.Errorf("repository.CreateEmployee() error = %v, wantErrMessage %v", err, tt.wantErrMessage)
				}
				return
			}
			if got == nil {
				t.Errorf("repository.CreateEmployee() = nil, want %v", tt.want)
				return
			}
			if !reflect.DeepEqual(got.EmployeeMail, tt.want.EmployeeMail) ||
				!reflect.DeepEqual(got.EmployeeUsername, tt.want.EmployeeUsername) ||
				!reflect.DeepEqual(got.EmployeePassword, tt.want.EmployeePassword) ||
				!reflect.DeepEqual(got.EmployeePhoneNumber, tt.want.EmployeePhoneNumber) ||
				!reflect.DeepEqual(got.Position, tt.want.Position) {
				t.Errorf("repository.CreateEmployee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_AuthenticateUser(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           *entities.Employee
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "Authenticate Succesfully",
			args: args{
				email:    "employeemail@mail.com",
				password: "employeepassword",
			},
			want: &entities.Employee{
				EmployeeMail:        "employeemail@mail.com",
				EmployeeUsername:    "employeeusername",
				EmployeePhoneNumber: "42",
				Position:            "tester",
				EmployeePassword:    "employeepassword",
			},
			wantErr: false,
		},
		{
			name: "Wrong Password",
			args: args{
				email:    "employeemail@mail.com",
				password: "wrongpassword",
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "username password combination is wrong",
		},
		{
			name: "Empty Email",
			args: args{
				email:    "",
				password: "password",
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "email cannot be null",
		},
		{
			name: "Empty Password",
			args: args{
				email:    "employeemail@mail.com",
				password: "",
			},
			want:           nil,
			wantErr:        true,
			wantErrMessage: "password cannot be null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Authenticate Succesfully":
				if err := testutils.SetupTestDataAuthenticateEmployee(dbConnection, tt.args.email, "employeeusername", "42", "tester", "employeepassword"); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			case "Wrong Password":
				if err := testutils.SetupTestDataAuthenticateEmployee(dbConnection, tt.args.email, "username", "42", "tester", "wR0nGpaẞw0rd"); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}

			defer testutils.CleanupTestDataAuthenticateEmployee(dbConnection, tt.args.email)

			r := &repository{
				DB: dbConnection,
			}
			got, err := r.AuthenticateUser(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && tt.want != nil {
				t.Errorf("repository.AuthenticateUser() = nil, want %v", tt.want)
				return
			}
			if tt.wantErr {
				if err == nil || err.Error() != tt.wantErrMessage {
					t.Errorf("repository.CreateEmployee() error = %v, wantErrMessage %v", err, tt.wantErrMessage)
				}
				return
			}
			if !reflect.DeepEqual(got.EmployeeMail, tt.want.EmployeeMail) ||
				!reflect.DeepEqual(got.EmployeeUsername, tt.want.EmployeeUsername) ||
				!reflect.DeepEqual(got.EmployeePassword, tt.want.EmployeePassword) ||
				!reflect.DeepEqual(got.EmployeePhoneNumber, tt.want.EmployeePhoneNumber) ||
				!reflect.DeepEqual(got.Position, tt.want.Position) {
				t.Errorf("repository.AuthenticateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_DeleteEmployee(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		employeeID string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "Delete Succesfully",
			args: args{
				employeeID: "testID",
			},
			wantErr: false,
		},
		{
			name: "User Not Found",
			args: args{
				employeeID: "testID",
			},
			wantErr:        true,
			wantErrMessage: "user not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.name == "Delete Succesfully" {
				if err := testutils.SetupTestDataDeleteEmployee(dbConnection, tt.args.employeeID); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				defer testutils.CleanupTestDataDeleteEmployee(dbConnection, tt.args.employeeID)

			}
			r := &repository{
				DB: dbConnection,
			}
			if err := r.DeleteEmployee(tt.args.employeeID); (err != nil) != tt.wantErr {
				t.Errorf("repository.DeleteEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
