package domain

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Isbn   string `json:"isbn"`
	Stock  int    `json:"stock"`
}

type BookUsecase interface {
	GetAllBooks() ([]*Book, error)
	GetBookId(isbn string) (*Book, error)
	CreateBook(book *Book) (*Book, error)
	DeleteBook(isbn string) error
}

type BookRepository interface {
	FindAll() ([]*Book, error)
	FindByIsbn(isbn string) (*Book, error)
	Create(book *Book) (*Book, error)
	Delete(isbn string) error
	UpdateStock(book *Book) error
}
