package employee

import (
	"errors"
	"reflect"
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateEmployee(employee *entities.Employee) (*entities.Employee, error) {
	args := m.Called(employee)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Employee), args.Error(1)
}

func (m *MockRepository) AuthenticateUser(email string, password string) (*entities.Employee, error) {
	args := m.Called(email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Employee), args.Error(1)
}

func (m *MockRepository) DeleteEmployee(employeeID string) error {
	args := m.Called(employeeID)
	return args.Error(0)
}

func Test_service_InsertEmployee(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		employee *entities.Employee
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Employee
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Inserted Employee Succcessfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				employee: &entities.Employee{
					EmployeeID:          "employee_id",
					EmployeeMail:        "employee@mail.com",
					EmployeeUsername:    "employee_username",
					EmployeePassword:    "employee_password",
					EmployeePhoneNumber: "employee_phone_number",
					Position:            "position",
				},
			},
			want: &entities.Employee{
				EmployeeID:          "employee_id",
				EmployeeMail:        "employee@mail.com",
				EmployeeUsername:    "employee_username",
				EmployeePassword:    "employee_password",
				EmployeePhoneNumber: "employee_phone_number",
				Position:            "position",
			},
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("CreateEmployee", mock.Anything).Return(&entities.Employee{
					EmployeeID:          "employee_id",
					EmployeeMail:        "employee@mail.com",
					EmployeeUsername:    "employee_username",
					EmployeePassword:    "employee_password",
					EmployeePhoneNumber: "employee_phone_number",
					Position:            "position",
				}, nil)
			},
		},
		{
			name: "Invalid Email",
			fields: fields{
				repo: nil,
			},
			args: args{
				employee: &entities.Employee{
					EmployeeID:          "employee_id",
					EmployeeMail:        "invalid",
					EmployeeUsername:    "employee_username",
					EmployeePassword:    "employee_password",
					EmployeePhoneNumber: "employee_phone_number",
					Position:            "position",
				},
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("CreateEmployee", mock.Anything).Return(nil, errors.New("invalid email address"))

			},
		},
		{
			name: "Username Too Short",
			fields: fields{
				repo: nil,
			},
			args: args{
				employee: &entities.Employee{
					EmployeeID:          "employee_id",
					EmployeeMail:        "employee_mail",
					EmployeeUsername:    "short",
					EmployeePassword:    "employee_password",
					EmployeePhoneNumber: "employee_phone_number",
					Position:            "position",
				},
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("CreateEmployee", mock.Anything).Return(nil, errors.New("username should be minimum 8 characters"))
			},
		},
		{
			name: "Username Too Long",
			fields: fields{
				repo: nil,
			},
			args: args{
				employee: &entities.Employee{
					EmployeeID:          "employee_id",
					EmployeeMail:        "employee_mail",
					EmployeeUsername:    "toolongtoolongtoolongtoolongtoolongtoolongtoolongtoolong",
					EmployeePassword:    "employee_password",
					EmployeePhoneNumber: "employee_phone_number",
					Position:            "position",
				},
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("CreateEmployee", mock.Anything).Return(nil, errors.New("username should be maximum 64 characters"))
			},
		},
		{
			name: "Password Too Short",
			fields: fields{
				repo: nil,
			},
			args: args{
				employee: &entities.Employee{
					EmployeeID:          "employee_id",
					EmployeeMail:        "employee_mail",
					EmployeeUsername:    "employee_username",
					EmployeePassword:    "short",
					EmployeePhoneNumber: "employee_phone_number",
					Position:            "position",
				},
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("CreateEmployee", mock.Anything).Return(nil, errors.New("password should be minimum 8 characters"))

			},
		},
		{
			name: "Password Too Long",
			fields: fields{
				repo: nil,
			},
			args: args{
				employee: &entities.Employee{
					EmployeeID:          "employee_id",
					EmployeeMail:        "employee_mail",
					EmployeeUsername:    "employee_username",
					EmployeePassword:    "toolongtoolongtoolongtoolongtoolongtoolongtoolongtoolong",
					EmployeePhoneNumber: "employee_phone_number",
					Position:            "position",
				},
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("CreateEmployee", mock.Anything).Return(nil, errors.New("password should be maximum 64 characters"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockSet(mockRepo)
			tt.fields.repo = mockRepo

			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.InsertEmployee(tt.args.employee)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.InsertEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.InsertEmployee() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_service_SignIn(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   *entities.Employee
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Signed Employee Successfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				email:    "test@example.com",
				password: "testpassword",
			},
			want: "test-token",
			want1: &entities.Employee{
				EmployeeID:          "employee_id",
				EmployeeMail:        "test@example.com",
				EmployeeUsername:    "employee_username",
				EmployeePassword:    "employee_password",
				EmployeePhoneNumber: "employee_phone_number",
				Position:            "position",
			},
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("AuthenticateUser", "test@example.com", "testpassword").Return(
					&entities.Employee{
						EmployeeID:          "employee_id",
						EmployeeMail:        "test@example.com",
						EmployeeUsername:    "employee_username",
						EmployeePassword:    "employee_password",
						EmployeePhoneNumber: "employee_phone_number",
						Position:            "position",
					}, nil)
			},
		},
		{
			name: "Employee Invalid Credentials",
			fields: fields{
				repo: nil,
			},
			args: args{
				email:    "test@example.com",
				password: "wrongpassword",
			},
			want:    "",
			want1:   nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("AuthenticateUser", "test@example.com", "wrongpassword").Return(
					nil, errors.New("invalid credentials"))
			},
		},
		{
			name: "Employee Empty Email",
			fields: fields{
				repo: nil,
			},
			args: args{
				email:    "",
				password: "testpassword",
			},
			want:    "",
			want1:   nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
			},
		},
		{
			name: "Employee Empty Password",
			fields: fields{
				repo: nil,
			},
			args: args{
				email:    "test@example.com",
				password: "",
			},
			want:    "",
			want1:   nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockSet(mockRepo)
			tt.fields.repo = mockRepo

			s := &service{
				repo: tt.fields.repo,
			}
			got, got1, err := s.SignIn(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != "" && tt.wantErr {
				t.Errorf("service.SignIn() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("service.SignIn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_service_DeleteEmployee(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		employeeID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Employee Deleted Successfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				employeeID: "test-id",
			},
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("DeleteEmployee", "test-id").Return(nil)
			},
		},
		{
			name: "Delete Empty Employee ID Error",
			fields: fields{
				repo: nil,
			},
			args: args{
				employeeID: "",
			},
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
			},
		},
		{
			name: "Delete Employee Error",
			fields: fields{
				repo: nil,
			},
			args: args{
				employeeID: "test-id",
			},
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("DeleteEmployee", "test-id").Return(errors.New("delete employee error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockRepository)
			tt.mockSet(mockRepo)
			tt.fields.repo = mockRepo

			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.DeleteEmployee(tt.args.employeeID); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
