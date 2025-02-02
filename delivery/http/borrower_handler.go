package http

import (
	"github.com/FRFebi/library-backend/domain"
	"github.com/gofiber/fiber/v2"
)

type BorrowerHandler struct {
	BorrowerUsecase domain.BorrowerUsecase
}

func NewBorrowerHandler(app fiber.Router, borrowerUsecase domain.BorrowerUsecase) {
	handler := &BorrowerHandler{BorrowerUsecase: borrowerUsecase}
	app.Get("/borrowers", handler.GetAllBorrowers)
	app.Get("/borrowers/:id", handler.GetBorrowerId)
	app.Post("/borrowers", handler.RegisterBorrower)
	app.Delete("/borrowers/:id", handler.UnregisterBorrower)
}

func (h *BorrowerHandler) GetAllBorrowers(c *fiber.Ctx) error {
	borrowers, err := h.BorrowerUsecase.GetAllBorrowers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(borrowers)
}

func (h *BorrowerHandler) GetBorrowerId(c *fiber.Ctx) error {
	borrowerId := c.Params("id")
	borrower, err := h.BorrowerUsecase.GetBorrowerId(borrowerId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(borrower)
}

func (h *BorrowerHandler) RegisterBorrower(c *fiber.Ctx) error {
	borrower := &domain.Borrower{}
	err := c.BodyParser(borrower)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	borrower, err = h.BorrowerUsecase.RegisterBorrower(borrower)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(borrower)
}

func (h *BorrowerHandler) UnregisterBorrower(c *fiber.Ctx) error {
	borrowerId := c.Params("id")
	err := h.BorrowerUsecase.UnregisterBorrower(borrowerId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(nil)
}
