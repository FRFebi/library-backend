package usecase

import (
	"time"

	"github.com/FRFebi/library-backend/domain"
)

type LoanUsecase struct {
	LoanRepository domain.LoanRepository
}

func NewLoanUsecase(loanRepository domain.LoanRepository) domain.LoanUsecase {
	return &LoanUsecase{LoanRepository: loanRepository}
}

func (u *LoanUsecase) GetAllLoans() ([]*domain.Loan, error) {
	return u.LoanRepository.FindAll()
}

func (u *LoanUsecase) GetLoanId(id int) (*domain.Loan, error) {
	return u.LoanRepository.FindById(id)
}

func (u *LoanUsecase) BorrowBook(loan *domain.Loan) (*domain.Loan, error) {
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
	now := time.Now()
	loan.ReturnDate.Time = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	loan.ReturnDate.Valid = true

	if loan.ReturnDate.Time.After(loan.DueDate) {
		loan.Status = "LATE"
	} else {
		loan.Status = "ON_TIME"
	}
	err = u.LoanRepository.Update(loan)
	if err != nil {
		return nil, err
	}
	return loan, err
}
