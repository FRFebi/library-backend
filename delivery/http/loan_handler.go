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

func (h *LoanHandler) GetAllLoans(c *fiber.Ctx) error {
	loans, err := h.LoanUsecase.GetAllLoans()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(loans)
}

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

func (h *LoanHandler) BorrowBook(c *fiber.Ctx) error {
	loan := &domain.Loan{}
	err := c.BodyParser(loan)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	loan, err = h.LoanUsecase.BorrowBook(loan)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(loan)
}

func (h *LoanHandler) ReturnBook(c *fiber.Ctx) error {
	loan := &domain.Loan{}
	err := c.BodyParser(loan)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	loan, err = h.LoanUsecase.ReturnBook(loan)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(loan)
}
