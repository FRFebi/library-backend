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

// GetAllBorrowers retrieves all borrowers
// @Summary Get all borrowers
// @Description Retrieve a list of all borrowers
// @Tags Borrowers
// @Accept json
// @Produce json
// @Success 200 {array} domain.Borrower
// @Failure 500
// @Router /api/borrowers [get]
func (h *BorrowerHandler) GetAllBorrowers(c *fiber.Ctx) error {
	borrowers, err := h.BorrowerUsecase.GetAllBorrowers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(borrowers)
}

// GetBorrowerId retrieves a borrower by ID
// @Summary Get borrower by ID
// @Description Retrieve a borrower by their unique ID
// @Tags Borrowers
// @Accept json
// @Produce json
// @Param id path string true "Borrower ID"
// @Success 200 {object} domain.Borrower
// @Failure 500
// @Router /api/borrowers/{id} [get]
func (h *BorrowerHandler) GetBorrowerId(c *fiber.Ctx) error {
	borrowerId := c.Params("id")
	borrower, err := h.BorrowerUsecase.GetBorrowerId(borrowerId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(borrower)
}

// RegisterBorrower registers a new borrower
// @Summary Register a new borrower
// @Description Create a new borrower in the system
// @Tags Borrowers
// @Accept json
// @Produce json
// @Param request body domain.Borrower true "Borrower Registration Request"
// @Success 200 {object} domain.Borrower
// @Failure 500
// @Router /api/borrowers [post]
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

// DeleteBorrower deletes a borrower by ID
// @Summary Delete a borrower
// @Description Remove a borrower from the system by their ID
// @Tags Borrowers
// @Accept json
// @Produce json
// @Param id path string true "Borrower ID"
// @Success 200 {object} nil
// @Failure 500
// @Router /api/borrowers/{id} [delete]
func (h *BorrowerHandler) UnregisterBorrower(c *fiber.Ctx) error {
	borrowerId := c.Params("id")
	err := h.BorrowerUsecase.UnregisterBorrower(borrowerId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(nil)
}
