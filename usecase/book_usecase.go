package usecase

import "github.com/FRFebi/library-backend/domain"

type BookUsecase struct {
	BookRepository domain.BookRepository
}

func NewBookUsecase(bookrepository domain.BookRepository) domain.BookUsecase {
	return &BookUsecase{BookRepository: bookrepository}
}

func (u *BookUsecase) GetAllBooks() ([]*domain.Book, error) {
	books, err := u.BookRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *BookUsecase) GetBookId(isbn string) (*domain.Book, error) {
	book, err := u.BookRepository.FindByIsbn(isbn)
	if err != nil || book == nil {
		return nil, err
	}

	return book, nil
}

func (u *BookUsecase) CreateBook(book *domain.Book) (*domain.Book, error) {
	return u.BookRepository.Create(book)
}

func (u *BookUsecase) DeleteBook(isbn string) error {
	err := u.BookRepository.Delete(isbn)
	if err != nil {
		return err
	}

	return nil
}
