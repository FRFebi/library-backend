package repository

import (
	"database/sql"

	"github.com/FRFebi/library-backend/domain"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) domain.BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) FindAll() ([]*domain.Book, error) {
	books := []*domain.Book{}

	query := "SELECT id, title, author, isbn, stock FROM book"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Isbn, &book.Stock)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *BookRepository) FindByIsbn(isbn string) (*domain.Book, error) {
	book := &domain.Book{}
	query := `SELECT id, title, author, isbn, stock FROM book WHERE isbn = $1`
	err := r.DB.QueryRow(query, isbn).Scan(&book.Id, &book.Title, &book.Author, &book.Isbn, &book.Stock)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) Create(book domain.Book) error {
	query := `INSERT INTO book (title, author, isbn, stock) VALUES($1,$2,$3,$4) RETURNING id`
	return r.DB.QueryRow(query, book.Title, book.Author, book.Isbn, book.Stock).Scan(&book.Id)
}

func (r *BookRepository) Delete(isbn string) error {
	query := `DELETE FROM book WHERE isbn = $1`
	_, err := r.DB.Exec(query, isbn)
	return err
}
