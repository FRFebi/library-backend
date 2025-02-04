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

// GetAllBooks retrieves all books
// @Summary Get all books
// @Description Retrieve a list of all books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} domain.Book
// @Failure 500
// @Router /api/books [get]
func (h *BookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.BookUsecase.GetAllBooks()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

// GetBookByID retrieves a book by its ISBN
// @Summary Get a book by ISBN
// @Description Retrieve a book by its ISBN
// @Tags Books
// @Accept json
// @Produce json
// @Param isbn path string true "Book ISBN"
// @Success 200 {object} domain.Book
// @Failure 500
// @Router /api/books/{isbn} [get]
func (h *BookHandler) GetBookId(c *fiber.Ctx) error {
	bookIsbn := c.Params("isbn")

	book, err := h.BookUsecase.GetBookId(bookIsbn)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Isbn   string `json:"isbn"`
	Stock  int    `json:"stock"`
}

// CreateBookRequest represents the request body for creating a new book.
// @Summary Create a new book
// @Description Create a new book with the provided details
// @Tags Books
// @Accept json
// @Produce json
// @Param request body CreateBookRequest true "Create Book Request"
// @Success 200 {object} domain.Book "Successfully created book"
// @Failure 500
// @Router /api/books [post]
func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	req := &CreateBookRequest{}
	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	book := &domain.Book{
		Title:  req.Title,
		Author: req.Author,
		Isbn:   req.Isbn,
		Stock:  req.Stock,
	}
	book, err := h.BookUsecase.CreateBook(book)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

// DeleteBook deletes a book by its ISBN
// @Summary Delete a book
// @Description Deletes a book from the library using its ISBN
// @Tags Books
// @Accept json
// @Produce json
// @Param isbn path string true "Book ISBN"
// @Success 200 {object} nil
// @Failure 500
// @Router /api/books/{isbn} [delete]
func (h *BookHandler) DeleteBook(c *fiber.Ctx) error {
	bookIsbn := c.Params("isbn")

	err := h.BookUsecase.DeleteBook(bookIsbn)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(nil)
}
