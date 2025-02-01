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
	GetBookId(id int) (*Book, error)
	CreateBook(book Book) error
	DeleteBook(id int) error
}

type BookRepository interface {
	FindAll() ([]*Book, error)
	FindById(id int) (*Book, error)
	Create(book Book) error
	Delete(id int) error
}
