package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/FRFebi/library-backend/domain"
)

type LoanUsecase struct {
	LoanRepository domain.LoanRepository
	BookRepository domain.BookRepository
}

func NewLoanUsecase(loanRepository domain.LoanRepository, bookRepository domain.BookRepository) domain.LoanUsecase {
	return &LoanUsecase{LoanRepository: loanRepository, BookRepository: bookRepository}
}

func (u *LoanUsecase) GetAllLoans() ([]*domain.Loan, error) {
	return u.LoanRepository.FindAll()
}

func (u *LoanUsecase) GetLoanId(id int) (*domain.Loan, error) {
	return u.LoanRepository.FindById(id)
}

func (u *LoanUsecase) BorrowBook(loan *domain.Loan) (*domain.Loan, error) {
	book, err := u.BookRepository.FindByIsbn(loan.BookIsbn)
	if err != nil {
		return nil, err
	}

	if book.Stock == 0 {
		return nil, fmt.Errorf("book is not available")
	}
	book.Stock--
	err = u.BookRepository.UpdateStock(book)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	loan.LoanDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	inWeek := time.Now().Add(time.Hour * 24 * 7)
	loan.DueDate = time.Date(inWeek.Year(), inWeek.Month(), inWeek.Day(), 0, 0, 0, 0, time.UTC)
	loan.Status = "NOT_RETURNED"
	return u.LoanRepository.Create(loan)
}

func (u *LoanUsecase) ReturnBook(loan *domain.Loan) (*domain.Loan, error) {
	loan, err := u.LoanRepository.FindById(loan.Id)
	if err != nil {
		return nil, err
	}
	if loan.Status != "NOT_RETURNED" {
		return nil, errors.New("book not already returned")
	}

	now := time.Now()
	loan.ReturnDate.Time = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	loan.ReturnDate.Valid = true

	if loan.ReturnDate.Time.After(loan.DueDate) {
		loan.Status = "LATE"
	} else {
		loan.Status = "ON_TIME"
	}
	book, err := u.BookRepository.FindByIsbn(loan.BookIsbn)
	if err != nil {
		return nil, err
	}
	book.Stock++
	err = u.BookRepository.UpdateStock(book)
	if err != nil {
		return nil, err
	}

	err = u.LoanRepository.Update(loan)
	if err != nil {
		return nil, err
	}
	return loan, err
}
