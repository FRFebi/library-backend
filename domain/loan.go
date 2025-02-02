package domain

import (
	"database/sql"
	"time"
)

type Loan struct {
	Id         int          `json:"id"`
	BookIsbn   string       `json:"book_isbn"`
	BorrowerId int          `json:"borrower_id"`
	LoanDate   time.Time    `json:"loan_date"`
	DueDate    time.Time    `json:"due_date"`
	ReturnDate sql.NullTime `json:"return_date"`
	Status     string       `json:"status"`
}

type LoanUsecase interface {
	GetAllLoans() ([]*Loan, error)
	GetLoanId(id int) (*Loan, error)
	BorrowBook(loan *Loan) (*Loan, error)
	ReturnBook(loan *Loan) (*Loan, error)
}

type LoanRepository interface {
	FindAll() ([]*Loan, error)
	FindById(id int) (*Loan, error)
	Create(loan *Loan) (*Loan, error)
	Update(loan *Loan) error
}
