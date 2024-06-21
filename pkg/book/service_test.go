package book

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

func (m *MockRepository) CreateBook(book *entities.Book) (*entities.Book, error) {
	args := m.Called(book)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Book), args.Error(1)
}

func (m *MockRepository) DeleteBook(bookID string) error {
	args := m.Called(bookID)
	return args.Error(0)
}

func (m *MockRepository) GetBook(bookID string) (*entities.Book, error) {
	args := m.Called(bookID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Book), args.Error(1)
}

func (m *MockRepository) GetBooks() ([]*entities.Book, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Book), args.Error(1)
}

func (m *MockRepository) GetBooksByID(bookIDList []string) ([]*entities.Book, error) {
	args := m.Called(bookIDList)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entities.Book), args.Error(1)
}

func Test_service_InsertBook(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		book *entities.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Book
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Inserted Book Successfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				book: testutils.ExampleBook,
			},
			want:    testutils.ExampleBook,
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("CreateBook", testutils.ExampleBook).Return(testutils.ExampleBook, nil)
			},
		},
		{
			name: "Insert Book Error",
			fields: fields{
				repo: nil,
			},
			args: args{
				book: testutils.ExampleBook,
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("CreateBook", testutils.ExampleBook).Return(nil, errors.New("error inserting book"))
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
			got, err := s.InsertBook(tt.args.book)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.InsertBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.InsertBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteBook(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		bookID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Book Deleted Successfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				bookID: "test-book-id",
			},
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("DeleteBook", "test-book-id").Return(nil)
			},
		},
		{
			name: "Empty Book ID",
			fields: fields{
				repo: nil,
			},
			args: args{
				bookID: "",
			},
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
			},
		},
		{
			name: "Delete Book Error",
			fields: fields{
				repo: nil,
			},
			args: args{
				bookID: "test-book-id",
			},
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("DeleteBook", "test-book-id").Return(errors.New("delete book error"))
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
			if err := s.DeleteBook(tt.args.bookID); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetBook(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		bookID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Book
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Fetched Book Successfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				bookID: "test-book-id",
			},
			want:    testutils.ExampleBook,
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("GetBook", "test-book-id").Return(testutils.ExampleBook, nil)
			},
		},
		{
			name: "Empty BookID",
			fields: fields{
				repo: nil,
			},
			args: args{
				bookID: "",
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
			},
		},
		{
			name: "Get Book Error",
			fields: fields{
				repo: nil,
			},
			args: args{
				bookID: "test-book-id",
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("GetBook", "test-book-id").Return(nil, errors.New("get book error"))
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
			got, err := s.GetBook(tt.args.bookID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetBooks(t *testing.T) {
	type fields struct {
		repo Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entities.Book
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Fetched Books Successfully",
			fields: fields{
				repo: nil,
			},
			want: []*entities.Book{
				testutils.ExampleBook,
			},
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("GetBooks").Return([]*entities.Book{testutils.ExampleBook}, nil)
			},
		},
		{
			name: "Get Books Error",
			fields: fields{
				repo: nil,
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("GetBooks").Return(nil, errors.New("get books error"))
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
			got, err := s.GetBooks()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetBooksByID(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		bookIdList []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Book
		wantErr bool
		mockSet func(mockRepo *MockRepository)
	}{
		{
			name: "Fetched Books By ID Successfully",
			fields: fields{
				repo: nil,
			},
			args: args{
				bookIdList: []string{"test-book-id-1", "test-book-id-2"},
			},
			want: []*entities.Book{
				testutils.ExampleBook,
				testutils.ExampleBook,
			},
			wantErr: false,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("GetBooksByID", []string{"test-book-id-1", "test-book-id-2"}).Return([]*entities.Book{testutils.ExampleBook, testutils.ExampleBook}, nil)
			},
		},
		{
			name: "Get Books By ID xError",
			fields: fields{
				repo: nil,
			},
			args: args{
				bookIdList: []string{"test-book-id-1", "test-book-id-2"},
			},
			want:    nil,
			wantErr: true,
			mockSet: func(mockRepo *MockRepository) {
				mockRepo.On("GetBooksByID", []string{"test-book-id-1", "test-book-id-2"}).Return(nil, errors.New("get books by ID error"))
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
			got, err := s.GetBooksByID(tt.args.bookIdList)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetBooksByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetBooksByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
