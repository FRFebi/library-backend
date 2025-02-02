package usecase

import "github.com/FRFebi/library-backend/domain"

type BorrowerUsecase struct {
	BorrowerRepository domain.BorrowerRepository
}

func NewBorrowerUsecase(BorrowerRepository domain.BorrowerRepository) domain.BorrowerUsecase {
	return &BorrowerUsecase{BorrowerRepository: BorrowerRepository}
}

func (u *BorrowerUsecase) GetAllBorrowers() ([]*domain.Borrower, error) {
	return u.BorrowerRepository.FindAll()
}

func (u *BorrowerUsecase) GetBorrowerId(id string) (*domain.Borrower, error) {
	return u.BorrowerRepository.FindById(id)
}

func (u *BorrowerUsecase) RegisterBorrower(borrower *domain.Borrower) (*domain.Borrower, error) {
	return u.BorrowerRepository.Create(borrower)
}

func (u *BorrowerUsecase) UnregisterBorrower(id string) error {
	return u.BorrowerRepository.Delete(id)
}
