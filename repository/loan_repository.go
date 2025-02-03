package repository

import (
	"database/sql"

	"github.com/FRFebi/library-backend/domain"
)

type LoanRepository struct {
	DB *sql.DB
}

func NewLoanRepository(db *sql.DB) domain.LoanRepository {
	return &LoanRepository{DB: db}
}

func (r *LoanRepository) FindAll() ([]*domain.Loan, error) {
	loans := []*domain.Loan{}
	query := "SELECT id, book_isbn, borrower_id, loan_date, due_date, return_date, status FROM loan"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		loan := domain.Loan{}
		err := rows.Scan(&loan.Id, &loan.BookIsbn, &loan.BorrowerId, &loan.LoanDate, &loan.DueDate, &loan.ReturnDate, &loan.Status)
		if err != nil {
			return nil, err
		}
		loans = append(loans, &loan)
	}
	if err != nil {
		return nil, err
	}
	return loans, nil
}

func (r *LoanRepository) FindById(id int) (*domain.Loan, error) {
	loan := domain.Loan{}
	query := "SELECT id, book_isbn, borrower_id, loan_date, due_date, return_date, status FROM loan WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&loan.Id, &loan.BookIsbn, &loan.BorrowerId, &loan.LoanDate, &loan.DueDate, &loan.ReturnDate, &loan.Status)
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (r *LoanRepository) Create(loan *domain.Loan) (*domain.Loan, error) {
	query := "INSERT INTO loan (book_isbn, borrower_id, loan_date, due_date, status) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	err := r.DB.QueryRow(query, loan.BookIsbn, loan.BorrowerId, loan.LoanDate, loan.DueDate, loan.Status).Scan(&loan.Id)
	return loan, err
}

func (r *LoanRepository) Update(loan *domain.Loan) error {
	query := "UPDATE loan SET return_date = $1, status = $2 WHERE id = $3"
	_, err := r.DB.Exec(query, loan.ReturnDate.Time, loan.Status, loan.Id)
	return err
}
