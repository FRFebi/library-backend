package domain

type Borrower struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type BorrowerUsecase interface {
	GetAllBorrowers() ([]*Borrower, error)
	GetBorrowerId(id string) (*Borrower, error)
	RegisterBorrower(borrower *Borrower) (*Borrower, error)
	UnregisterBorrower(id string) error
}

type BorrowerRepository interface {
	FindAll() ([]*Borrower, error)
	FindById(id string) (*Borrower, error)
	Create(borrower *Borrower) (*Borrower, error)
	Delete(id string) error
}
