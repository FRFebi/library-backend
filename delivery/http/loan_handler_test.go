package http

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/FRFebi/library-backend/domain"
	"github.com/FRFebi/library-backend/infrastructure"
	"github.com/FRFebi/library-backend/repository"
	"github.com/FRFebi/library-backend/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func LoanSetup(db *sql.DB) domain.LoanUsecase {

	loanRepo := repository.NewLoanRepository(db)
	bookRepo := repository.NewBookRepository(db)
	return usecase.NewLoanUsecase(loanRepo, bookRepo)

}

func TestLoanHandler(t *testing.T) {
	app := fiber.New()
	api := app.Group("/api")
	db := infrastructure.NewPostgreeDB()
	bookUC := BookSetup(db)
	borrowerUC := BorrowSetup(db)
	loanUC := LoanSetup(db)

	NewLoanHandler(api, loanUC)
	NewBookHandler(api, bookUC)
	NewBorrowerHandler(api, borrowerUC)

	book := &domain.Book{}
	t.Run("Get All Books", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseBook []*domain.Book
		json.NewDecoder(resp.Body).Decode(&responseBook)
		assert.NotEqual(t, 0, len(responseBook))

		book = responseBook[0]
	})

	borrower := &domain.Borrower{}
	t.Run("Get All Borrowers", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/borrowers", nil)
		resp, err := app.Test(req)

		assert.Nil(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseBorrower []*domain.Borrower
		json.NewDecoder(resp.Body).Decode(&responseBorrower)
		assert.NotEqual(t, 0, len(responseBorrower))

		borrower = responseBorrower[0]
	})

	loan := &domain.Loan{BookIsbn: book.Isbn, BorrowerId: borrower.Id}
	t.Run("Borrow Book", func(t *testing.T) {
		jsonLoan, _ := json.Marshal(loan)

		req := httptest.NewRequest(http.MethodPost, "/api/loans/borrow", bytes.NewBuffer(jsonLoan))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseLoan *domain.Loan
		json.NewDecoder(resp.Body).Decode(&responseLoan)
		loan.Id = responseLoan.Id
		book.Stock--
	})

	t.Run("Get Book by ISBN", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/books/"+book.Isbn, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseBook *domain.Book
		json.NewDecoder(resp.Body).Decode(&responseBook)
		assert.Equal(t, book.Stock, responseBook.Stock)
	})

	loanId := strconv.Itoa(loan.Id)
	t.Run("Get Loan by ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/loans/"+loanId, nil)
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseLoan *domain.Loan
		json.NewDecoder(resp.Body).Decode(&responseLoan)
		assert.Equal(t, "NOT_RETURNED", responseLoan.Status)
	})

	t.Run("Return Book", func(t *testing.T) {
		jsonLoan, _ := json.Marshal(loan)

		req := httptest.NewRequest(http.MethodPost, "/api/loans/return", bytes.NewBuffer(jsonLoan))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

	})

	t.Run("Get Loan by ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/loans/"+loanId, nil)
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseLoan *domain.Loan
		json.NewDecoder(resp.Body).Decode(&responseLoan)
		assert.Equal(t, "ON_TIME", responseLoan.Status)
		book.Stock++
	})

	t.Run("Get Book by ISBN", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/books/"+book.Isbn, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseBook *domain.Book
		json.NewDecoder(resp.Body).Decode(&responseBook)
		assert.Equal(t, book.Stock, responseBook.Stock)
	})
	t.Run("Get All Loans", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/loans", nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
