package repository

import (
	"database/sql"

	"github.com/FRFebi/library-backend/domain"
)

type BorrowerRepository struct {
	DB *sql.DB
}

func NewBorrowerRepository(db *sql.DB) domain.BorrowerRepository {
	return &BorrowerRepository{DB: db}
}

func (r *BorrowerRepository) FindAll() ([]*domain.Borrower, error) {
	borrowers := []*domain.Borrower{}
	query := "SELECT id, name, email, phone FROM borrower"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		borrower := domain.Borrower{}
		err = rows.Scan(&borrower.Id, &borrower.Name, &borrower.Email, &borrower.Phone)
		if err != nil {
			return nil, err
		}
		borrowers = append(borrowers, &borrower)
	}

	return borrowers, nil
}

func (r *BorrowerRepository) FindById(id string) (*domain.Borrower, error) {
	borrower := &domain.Borrower{}
	query := "SELECT id, name, email, phone FROM borrower WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&borrower.Id, &borrower.Name, &borrower.Email, &borrower.Phone)
	if err != nil {
		return nil, err
	}

	return borrower, nil
}

func (r *BorrowerRepository) Create(borrower *domain.Borrower) (*domain.Borrower, error) {
	query := "INSERT INTO borrower(name, email, phone) VALUES ($1, $2, $3) RETURNING id"
	err := r.DB.QueryRow(query, &borrower.Name, &borrower.Email, &borrower.Phone).Scan(&borrower.Id)
	if err != nil {
		return nil, err
	}
	return borrower, nil
}

func (r *BorrowerRepository) Delete(id string) error {
	query := "DELETE FROM borrower WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
