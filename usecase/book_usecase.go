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

func (u *BookUsecase) GetBookId(id int) (*domain.Book, error) {
	book, err := u.BookRepository.FindById(id)
	if err != nil || book == nil {
		return nil, err
	}

	return book, nil
}

func (u *BookUsecase) CreateBook(book domain.Book) error {
	err := u.BookRepository.Create(book)
	if err != nil {
		return err
	}

	return nil
}

func (u *BookUsecase) DeleteBook(id int) error {
	err := u.BookRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
