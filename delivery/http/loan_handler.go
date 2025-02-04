package http

import (
	"strconv"

	"github.com/FRFebi/library-backend/domain"
	"github.com/gofiber/fiber/v2"
)

type LoanHandler struct {
	LoanUsecase domain.LoanUsecase
}

func NewLoanHandler(app fiber.Router, loanUsecase domain.LoanUsecase) {
	handler := &LoanHandler{LoanUsecase: loanUsecase}
	app.Get("/loans", handler.GetAllLoans)
	app.Get("/loans/:id", handler.GetLoanId)
	app.Post("/loans/borrow", handler.BorrowBook)
	app.Post("/loans/return", handler.ReturnBook)
}

// GetAllLoans retrieves all loans
// @Summary Get all loans
// @Description Retrieve a list of all loans
// @Tags Loans
// @Accept json
// @Produce json
// @Success 200 {array} domain.Loan  "List of loans"
// @Failure 500
// @Router /api/loans [get]
func (h *LoanHandler) GetAllLoans(c *fiber.Ctx) error {
	loans, err := h.LoanUsecase.GetAllLoans()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(loans)
}

// GetLoanId retrieves a loan by ID
// @Summary Get loan by ID
// @Description Retrieve a loan by its ID
// @Tags Loans
// @Accept json
// @Produce json
// @Param id path string true "Loan ID"
// @Success 200 {object} domain.Loan "Loan details"
// @Failure 500
// @Router /api/loans/{id} [get]
func (h *LoanHandler) GetLoanId(c *fiber.Ctx) error {
	loanId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	loan, err := h.LoanUsecase.GetLoanId(loanId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(loan)
}

type BorrowBook struct {
	BookIsbn   string `json:"book_isbn"`
	BorrowerId int    `json:"borrower_id"`
}

// BorrowBook allows a borrower to borrow a book
// @Summary Borrow a book
// @Description Borrow a book by providing the loan details
// @Tags Loans
// @Accept json
// @Produce json
// @Param request body BorrowBook true "Borrow book request"
// @Success 200 {object} domain.Loan "Loan details"
// @Failure 500
// @Router /api/loans/borrow [post]
func (h *LoanHandler) BorrowBook(c *fiber.Ctx) error {
	req := &BorrowBook{}
	err := c.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	loan := &domain.Loan{BookIsbn: req.BookIsbn, BorrowerId: req.BorrowerId}
	loan, err = h.LoanUsecase.BorrowBook(loan)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(loan)
}

type ReturnBook struct {
	Id int `json:"id"`
}

// ReturnBook allows a borrower to return a book
// @Summary Return a book
// @Description Return a borrowed book
// @Tags Loans
// @Accept json
// @Produce json
// @Param request body ReturnBook true "Return book request"
// @Success 200
// @Failure 500
// @Router /api/loans/return [post]
func (h *LoanHandler) ReturnBook(c *fiber.Ctx) error {
	req := &ReturnBook{}
	err := c.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	loan := &domain.Loan{Id: req.Id}
	loan, err = h.LoanUsecase.ReturnBook(loan)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(loan)
}
