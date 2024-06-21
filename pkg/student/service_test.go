package student

import (
	"errors"
	"reflect"
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) AddStudent(student *entities.Student) (*entities.Student, error) {
	args := m.Called(student)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Student), args.Error(1)
}

func (m *MockRepository) AuthenticateStudent(email, password string) (*entities.Student, error) {
	args := m.Called(email, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Student), args.Error(1)
}

func (m *MockRepository) GetStudentByEmail(email string) (*entities.Student, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Student), args.Error(1)
}

func (m *MockRepository) UpdateStudentPassword(studentID string, hashedPassword string) error {
	args := m.Called(studentID, hashedPassword)
	return args.Error(0)
}

func (m *MockRepository) BorrowBook(bookID, studentID string) (string, error) {
	args := m.Called(bookID, studentID)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *MockRepository) DeliverBook(borrowID, bookID, studentID string) (string, error) {
	args := m.Called(borrowID, bookID, studentID)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *MockRepository) ExtendDate(borrowID string) (string, error) {
	args := m.Called(borrowID)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

func (m *MockRepository) GetBorrowedBooks(studentID string) ([]entities.BorrowedBook, error) { // Remove the pointers
	args := m.Called(studentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.BorrowedBook), args.Error(1)
}

func Test_service_InsertStudent(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		student *entities.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Student
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Student Inserted Succesfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				student: testutils.ExampleStudent,
			},
			want:    testutils.ExampleStudent,
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("AddStudent", testutils.ExampleStudent).Return(testutils.ExampleStudent, nil)
			},
		},
		{
			name: "Insert Student Error",
			fields: fields{
				repo: nil,
			},
			args: args{
				student: testutils.ExampleStudent,
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("AddStudent", testutils.ExampleStudent).Return(nil, errors.New("insert student error"))
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
			got, err := s.InsertStudent(tt.args.student)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.InsertStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.InsertStudent() = %v, want %v", got, tt.want)
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
		want1   *entities.Student
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Signed Successfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				email:    "test@example.com",
				password: "testpassword",
			},
			want: "test-token",
			want1: &entities.Student{
				StudentID:   "test-student-id",
				StudentMail: "test@example.com",
			},
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("AuthenticateStudent", "test@example.com", "testpassword").Return(
					&entities.Student{
						StudentID:   "test-student-id",
						StudentMail: "test@example.com",
					}, nil)
			},
		},
		{
			name: "Invalid Credentials",
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
				mockRepo.On("AuthenticateStudent", "test@example.com", "wrongpassword").Return(
					nil, errors.New("invalid credentials"))
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
			if got != tt.want && tt.wantErr {
				t.Errorf("service.SignIn() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("service.SignIn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// TODO check one more time
// func Test_service_RequestPasswordReset(t *testing.T) {
// 	type fields struct {
// 		repo Repository
// 	}
// 	type args struct {
// 		email string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 		mockSet func(mockRepo *MockRepository)
// 	}{
// 		{
// 			name: "RequestPasswordResetSuccess",
// 			fields: fields{
// 				repo: nil,
// 			},
// 			args: args{
// 				email: "test@example.com",
// 			},
// 			wantErr: false,
// 			mockSet: func(mockRepo *MockRepository) {
// 				mockRepo.On("GetStudentByEmail", "test@example.com").Return(
// 					&entities.Student{
// 						StudentID:   "test-student-id",
// 						StudentMail: "test@example.com",
// 					}, nil)
// 			},
// 		},
// 		{
// 			name: "RequestPasswordResetError",
// 			fields: fields{
// 				repo: nil,
// 			},
// 			args: args{
// 				email: "test@example.com",
// 			},
// 			wantErr: true,
// 			mockSet: func(mockRepo *MockRepository) {
// 				mockRepo.On("GetStudentByEmail", "test@example.com").Return(nil, errors.New("get student by email error"))
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockRepo := new(MockRepository)
// 			tt.mockSet(mockRepo)
// 			tt.fields.repo = mockRepo
// 			s := &service{
// 				repo: tt.fields.repo,
// 			}
// 			if err := s.RequestPasswordReset(tt.args.email); (err != nil) != tt.wantErr {
// 				t.Errorf("service.RequestPasswordReset() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func Test_service_ResetPassword(t *testing.T) {
// 	type fields struct {
// 		repo Repository
// 	}
// 	type args struct {
// 		token       string
// 		newPassword string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 		mockSet func(mockRepo *MockRepository)
// 	}{
// 		{
// 			name: "ResetPasswordSuccess",
// 			fields: fields{
// 				repo: nil,
// 			},
// 			args: args{
// 				token:       "testtoken",
// 				newPassword: "newpassword",
// 			},
// 			wantErr: false,
// 			mockSet: func(mockRepo *MockRepository) {
// 				mockRepo.On("UpdateStudentPassword", "test-student-id", "hashed-password").Return(nil)
// 			},
// 		},
// 		{
// 			name: "ResetPasswordError",
// 			fields: fields{
// 				repo: nil,
// 			},
// 			args: args{
// 				token:       "testtoken",
// 				newPassword: "newpassword",
// 			},
// 			wantErr: true,
// 			mockSet: func(mockRepo *MockRepository) {
// 				mockRepo.On("UpdateStudentPassword", "test-student-id", "hashed-password").Return(errors.New("update student password error"))
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockRepo := new(MockRepository)
// 			tt.mockSet(mockRepo)
// 			tt.fields.repo = mockRepo
// 			s := &service{
// 				repo: tt.fields.repo,
// 			}
// 			if err := s.ResetPassword(tt.args.token, tt.args.newPassword); (err != nil) == tt.wantErr {
// 				t.Errorf("service.ResetPassword() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func Test_service_BookBorrow(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		bookID    string
		studentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Book Borrowed Succesfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				bookID:    "test-book-id",
				studentID: "test-student-id",
			},
			want:    "test-borrow-id",
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("BorrowBook", "test-book-id", "test-student-id").Return("test-borrow-id", nil)
			},
		},
		{
			name: "Book Borrow Error",
			fields: fields{
				repo: nil,
			},
			args: args{
				bookID:    "test-book-id",
				studentID: "test-student-id",
			},
			want:    "",
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("BorrowBook", "test-book-id", "test-student-id").Return("", errors.New("borrow book error"))
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
			got, err := s.BookBorrow(tt.args.bookID, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.BookBorrow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.BookBorrow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeliverBook(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		borrowId  string
		bookId    string
		studentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Book Delivered Successfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				borrowId:  "test-borrow-id",
				bookId:    "test-book-id",
				studentID: "test-student-id",
			},
			want:    "Book delivered successfully",
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("DeliverBook", "test-borrow-id", "test-book-id", "test-student-id").Return("Book delivered successfully", nil)
			},
		},
		{
			name: "Deliver Book Error",
			fields: fields{
				repo: nil,
			},
			args: args{
				borrowId:  "test-borrow-id",
				bookId:    "test-book-id",
				studentID: "test-student-id",
			},
			want:    "",
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("DeliverBook", "test-borrow-id", "test-book-id", "test-student-id").Return("", errors.New("deliver book error"))
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
			got, err := s.DeliverBook(tt.args.borrowId, tt.args.bookId, tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.DeliverBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.DeliverBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_ExtendDate(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		borrowID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Date Extended Successfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				borrowID: "test-borrow-id",
			},
			want:    "Book extend date successfully",
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("ExtendDate", "test-borrow-id").Return("Book extend date successfully", nil)
			},
		},
		{
			name: "Extend Date Error",
			fields: fields{
				repo: nil,
			},
			args: args{
				borrowID: "test-borrow-id",
			},
			want:    "",
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("ExtendDate", "test-borrow-id").Return("", errors.New("extend book date error"))
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
			got, err := s.ExtendDate(tt.args.borrowID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.ExtendDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.ExtendDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetBorrowedBooks(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		studentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.BorrowedBook
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Books Fetched Successfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				studentID: "test-student-id",
			},
			want: []*entities.BorrowedBook{
				{
					BorrowID:  "test-borrow-id-1",
					BookID:    "test-book-id-1",
					StudentID: "test-student-id",
				},
				{
					BorrowID:  "test-borrow-id-2",
					BookID:    "test-book-id-2",
					StudentID: "test-student-id",
				},
			},
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("GetBorrowedBooks", "test-student-id").Return(
					[]entities.BorrowedBook{
						{
							BorrowID:  "test-borrow-id-1",
							BookID:    "test-book-id-1",
							StudentID: "test-student-id",
						},
						{
							BorrowID:  "test-borrow-id-2",
							BookID:    "test-book-id-2",
							StudentID: "test-student-id",
						},
					}, nil)
			},
		},
		{
			name: "Get Borrowed Books Error",
			fields: fields{
				repo: nil,
			},
			args: args{
				studentID: "test-student-id",
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("GetBorrowedBooks", "test-student-id").Return(nil, errors.New("get borrowed books error"))
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
			got, err := s.GetBorrowedBooks(tt.args.studentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetBorrowedBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetBorrowedBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}
