package book

import (
	"os"
	"reflect"
	"testing"

	"github.com/afurgapil/library-management-system/pkg/entities"
	testutils "github.com/afurgapil/library-management-system/pkg/testUtils"
	"github.com/jackc/pgx/v4"
)

// Book - DB Connection
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

func Test_repository_CreateBook(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		book *entities.Book
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           *entities.Book
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "Create Succesfully",
			args: args{
				book: &entities.Book{
					BookID:          "book_id",
					Title:           "book_title",
					Author:          "book_Author",
					Genre:           "book_genre",
					PublicationDate: "99.99.9999",
					Publisher:       "book_publisher",
					ISBN:            "book_isbn",
					PageCount:       20,
					ShelfNumber:     "book_shelf",
					Language:        "book_language",
					Donor:           "book_donor",
				},
			},
			fields: fields{
				DB: dbConnection,
			},
			want: &entities.Book{
				BookID:          "book_id",
				Title:           "book_title",
				Author:          "book_Author",
				Genre:           "book_genre",
				PublicationDate: "99.99.9999",
				Publisher:       "book_publisher",
				ISBN:            "book_isbn",
				PageCount:       20,
				ShelfNumber:     "book_shelf",
				Language:        "book_language",
				Donor:           "book_donor",
			},
			wantErr: false,
		},
		{
			name: "Duplicate BookID",
			args: args{
				book: &entities.Book{
					BookID:          "book_id",
					Title:           "book_title",
					Author:          "book_Author",
					Genre:           "book_genre",
					PublicationDate: "99.99.9999",
					Publisher:       "book_publisher",
					ISBN:            "book_isbn",
					PageCount:       20,
					ShelfNumber:     "book_shelf",
					Language:        "book_language",
					Donor:           "book_donor",
				},
			},
			fields: fields{
				DB: dbConnection,
			},
			want: &entities.Book{
				BookID:          "book_id",
				Title:           "book_title",
				Author:          "book_Author",
				Genre:           "book_genre",
				PublicationDate: "99.99.9999",
				Publisher:       "book_publisher",
				ISBN:            "book_isbn",
				PageCount:       20,
				ShelfNumber:     "book_shelf",
				Language:        "book_language",
				Donor:           "book_donor",
			},
			wantErr:        true,
			wantErrMessage: "duplicated ID",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Duplicate BookID":
				if err := testutils.SetupTestDataCreateBook(tt.fields.DB, tt.args.book); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			case "Duplicate ISBN":
				if err := testutils.SetupTestDataCreateBook(tt.fields.DB, tt.args.book); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}
			defer testutils.CleanupTestDataCreateBook(tt.fields.DB, tt.args.book.BookID)

			r := &repository{
				DB: tt.fields.DB,
			}
			got, err := r.CreateBook(tt.args.book)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if err == nil || err.Error() != tt.wantErrMessage {
					t.Errorf("repository.CreateBook() error = %v, wantErrMessage = %v", err, tt.wantErrMessage)
				}
				return
			}

			if !reflect.DeepEqual(got.Title, tt.want.Title) &&
				!reflect.DeepEqual(got.Author, tt.want.Author) &&
				!reflect.DeepEqual(got.Genre, tt.want.Genre) &&
				!reflect.DeepEqual(got.PublicationDate, tt.want.PublicationDate) &&
				!reflect.DeepEqual(got.Publisher, tt.want.Publisher) &&
				!reflect.DeepEqual(got.ISBN, tt.want.ISBN) &&
				!reflect.DeepEqual(got.PageCount, tt.want.PageCount) &&
				!reflect.DeepEqual(got.ShelfNumber, tt.want.ShelfNumber) &&
				!reflect.DeepEqual(got.Language, tt.want.Language) &&
				!reflect.DeepEqual(got.Donor, tt.want.Donor) {
				t.Errorf("repository.CreateBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_DeleteBook(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		bookID string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "Deleted Succesfully",
			args: args{
				bookID: "delete-book-ID",
			},
			fields: fields{
				DB: dbConnection,
			},
			wantErr: false,
		},
		{
			name: "Empty BookID",
			args: args{
				bookID: "",
			},
			fields: fields{
				DB: dbConnection,
			},
			wantErr:        true,
			wantErrMessage: "book ID cannot be empty",
		},
		{
			name: "Book Not Found",
			args: args{
				bookID: "delete-book-ID",
			},
			fields: fields{
				DB: dbConnection,
			},
			wantErr:        true,
			wantErrMessage: "book not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Deleted Succesfully":
				if err := testutils.SetupTestDataDeleteBook(dbConnection, tt.args.bookID); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
			}
			defer testutils.CleanupTestDataDeleteBook(dbConnection, "delete-book-ID")
			r := &repository{
				DB: tt.fields.DB,
			}
			if err := r.DeleteBook(tt.args.bookID); (err != nil) != tt.wantErr {
				t.Errorf("repository.DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_GetBook(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		bookID string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           *entities.Book
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "Get Book Succesfully",
			args: args{
				bookID: "get_book_id",
			},
			fields: fields{
				DB: dbConnection,
			},
			want: &entities.Book{
				BookID:          "get_book_id",
				Title:           "book_title",
				Author:          "book_Author",
				Genre:           "book_genre",
				PublicationDate: "99-99-9999",
				Publisher:       "book_publisher",
				ISBN:            "book_isbn",
				PageCount:       20,
				ShelfNumber:     "book_shelf",
				Language:        "book_language",
				Donor:           "book_donor",
			},
			wantErr: false,
		},
		{
			name: "Empty bookID",
			args: args{
				bookID: "",
			},
			fields: fields{
				DB: dbConnection,
			},
			wantErr:        true,
			wantErrMessage: "book ID cannot be empty",
		},
		{
			name: "Book Not Found",
			args: args{
				bookID: "book_not_found",
			},
			fields: fields{
				DB: dbConnection,
			},
			wantErr:        true,
			wantErrMessage: "book not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.name {
			case "Get Book Succesfully":
				if err := testutils.SetupTestDataGetBook(dbConnection, tt.args.bookID); err != nil {
					t.Fatalf("Failed to set up test data: %v", err)
				}
				defer testutils.CleanupTestDataGetBook(dbConnection, "get-book-ID")
			}
			defer testutils.CleanupTestDataGetBook(dbConnection, "get-book-ID")
			r := &repository{
				DB: tt.fields.DB,
			}
			got, err := r.GetBook(tt.args.bookID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetBook() = %v, want %v", got, tt.want)
			}
			if tt.name == "Get Book Succesfully" {
				if err := testutils.CleanupTestDataGetBook(dbConnection, tt.args.bookID); err != nil {
					t.Fatalf("Failed to clean up test data: %v", err)
				}
			}

		})
	}
}

func Test_repository_GetBooks(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entities.Book
		wantErr bool
	}{
		{
			name: "Get Books Successfully",
			fields: fields{
				DB: dbConnection,
			},
			want: []*entities.Book{
				{
					BookID:          "get_book_id1",
					Title:           "book_title1",
					Author:          "book_Author1",
					Genre:           "book_genre1",
					PublicationDate: "99-99-9999",
					Publisher:       "book_publisher1",
					ISBN:            "book_isbn1",
					PageCount:       20,
					ShelfNumber:     "book_shelf1",
					Language:        "book_language1",
					Donor:           "book_donor1",
				},
				{
					BookID:          "get_book_id2",
					Title:           "book_title2",
					Author:          "book_Author2",
					Genre:           "book_genre2",
					PublicationDate: "99-99-9999",
					Publisher:       "book_publisher2",
					ISBN:            "book_isbn2",
					PageCount:       20,
					ShelfNumber:     "book_shelf2",
					Language:        "book_language2",
					Donor:           "book_donor2",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testutils.CleanupTestDataGetBooks(dbConnection, tt.want); err != nil {
				t.Fatalf("Failed to clean up test data: %v", err)
			}

			if err := testutils.SetupTestDataGetBooks(dbConnection, tt.want); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}
			defer testutils.CleanupTestDataGetBooks(dbConnection, tt.want)

			r := &repository{
				DB: tt.fields.DB,
			}
			got, err := r.GetBooks()
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetBooksByID(t *testing.T) {
	type fields struct {
		DB *pgx.Conn
	}
	type args struct {
		bookIDList []string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           []*entities.Book
		wantErr        bool
		wantErrMessage string
	}{
		{
			name: "Get Books by ID succesfully",
			fields: fields{
				DB: dbConnection,
			},
			args: args{
				bookIDList: []string{"get_book_id1", "get_book_id2"},
			},
			want: []*entities.Book{
				{
					BookID:          "get_book_id1",
					Title:           "book_title1",
					Author:          "book_Author1",
					Genre:           "book_genre1",
					PublicationDate: "99-99-9999",
					Publisher:       "book_publisher1",
					ISBN:            "book_isbn1",
					PageCount:       20,
					ShelfNumber:     "book_shelf1",
					Language:        "book_language1",
					Donor:           "book_donor1",
				},
				{
					BookID:          "get_book_id2",
					Title:           "book_title2",
					Author:          "book_Author2",
					Genre:           "book_genre2",
					PublicationDate: "99-99-9999",
					Publisher:       "book_publisher2",
					ISBN:            "book_isbn2",
					PageCount:       20,
					ShelfNumber:     "book_shelf2",
					Language:        "book_language2",
					Donor:           "book_donor2",
				},
			},
		},
		{
			name: "Missing Book(s)",
			fields: fields{
				DB: dbConnection,
			},
			args: args{
				bookIDList: []string{"get_book_id1", "get_book_id2"},
			},
			want: []*entities.Book{
				{
					BookID:          "get_book_id1",
					Title:           "book_title1",
					Author:          "book_Author1",
					Genre:           "book_genre1",
					PublicationDate: "99-99-9999",
					Publisher:       "book_publisher1",
					ISBN:            "book_isbn1",
					PageCount:       20,
					ShelfNumber:     "book_shelf1",
					Language:        "book_language1",
					Donor:           "book_donor1",
				},
			},
			wantErr:        true,
			wantErrMessage: "number of returned books does not match expected",
		},
		{
			name: "Books Not Found",
			fields: fields{
				DB: dbConnection,
			},
			args: args{
				bookIDList: []string{"get_book_id1", "get_book_id2"},
			},
			want:           []*entities.Book{},
			wantErr:        true,
			wantErrMessage: "books not found",
		},
		{
			name: "Invalid Book ID(s)",
			fields: fields{
				DB: dbConnection,
			},
			args: args{
				bookIDList: []string{"get_book_id1", "get_book_id2", ""},
			},
			want: []*entities.Book{
				{
					BookID:          "get_book_id1",
					Title:           "book_title1",
					Author:          "book_Author1",
					Genre:           "book_genre1",
					PublicationDate: "99-99-9999",
					Publisher:       "book_publisher1",
					ISBN:            "book_isbn1",
					PageCount:       20,
					ShelfNumber:     "book_shelf1",
					Language:        "book_language1",
					Donor:           "book_donor1",
				},
				{
					BookID:          "get_book_id2",
					Title:           "book_title2",
					Author:          "book_Author2",
					Genre:           "book_genre2",
					PublicationDate: "99-99-9999",
					Publisher:       "book_publisher2",
					ISBN:            "book_isbn2",
					PageCount:       20,
					ShelfNumber:     "book_shelf2",
					Language:        "book_language2",
					Donor:           "book_donor2",
				},
				{
					BookID:          "get_book_id3",
					Title:           "book_title3",
					Author:          "book_Author3",
					Genre:           "book_genre3",
					PublicationDate: "99-99-9999",
					Publisher:       "book_publisher3",
					ISBN:            "book_isbn3",
					PageCount:       30,
					ShelfNumber:     "book_shelf3",
					Language:        "book_language3",
					Donor:           "book_donor3",
				},
			},
			wantErr:        true,
			wantErrMessage: "book ID list contains empty strings",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testutils.CleanupTestDataGetBooks(dbConnection, tt.want); err != nil {
				t.Fatalf("Failed to clean up test data: %v", err)
			}

			if err := testutils.SetupTestDataGetBooks(dbConnection, tt.want); err != nil {
				t.Fatalf("Failed to set up test data: %v", err)
			}
			defer testutils.CleanupTestDataGetBooks(dbConnection, tt.want)

			r := &repository{
				DB: tt.fields.DB,
			}
			got, err := r.GetBooksByID(tt.args.bookIDList)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetBooksByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err == nil || err.Error() != tt.wantErrMessage {
					t.Errorf("repository.CreateBook() error = %v, wantErrMessage = %v", err, tt.wantErrMessage)
				}
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetBooksByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
