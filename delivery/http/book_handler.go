package http

import (
	"github.com/FRFebi/library-backend/domain"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	BookUsecase domain.BookUsecase
}

func NewBookHandler(app fiber.Router, bookUsecase domain.BookUsecase) {
	handler := &BookHandler{BookUsecase: bookUsecase}
	app.Get("/books", handler.GetAllBooks)
	app.Get("/books/:isbn", handler.GetBookId)
	app.Post("/books", handler.CreateBook)
	app.Delete("/books/:isbn", handler.DeleteBook)
}

func (h *BookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.BookUsecase.GetAllBooks()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func (h *BookHandler) GetBookId(c *fiber.Ctx) error {
	bookIsbn := c.Params("isbn")

	book, err := h.BookUsecase.GetBookId(bookIsbn)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	book := &domain.Book{}
	if err := c.BodyParser(book); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	book, err := h.BookUsecase.CreateBook(book)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	bookIsbn := c.Params("isbn")

	err := h.BookUsecase.DeleteBook(bookIsbn)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(nil)
}
