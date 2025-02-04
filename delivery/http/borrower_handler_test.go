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

func BorrowSetup(db *sql.DB) domain.BorrowerUsecase {
	borrowerRepo := repository.NewBorrowerRepository(db)
	return usecase.NewBorrowerUsecase(borrowerRepo)
}

func TestBorrowerHandler(t *testing.T) {
	app := fiber.New()
	api := app.Group("/api")
	db := infrastructure.NewPostgreeDB()

	borrowerUC := BorrowSetup(db)
	NewBorrowerHandler(api, borrowerUC)

	newBorrower := &domain.Borrower{Name: "Febi", Email: "febi@mail.com", Phone: "123456"}
	t.Run("Register Borrower", func(t *testing.T) {
		jsonBorrower, _ := json.Marshal(newBorrower)

		// Act
		req := httptest.NewRequest(http.MethodPost, "/api/borrowers", bytes.NewBuffer(jsonBorrower))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		assert.Nil(t, err)

		// Assert
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseBorrower *domain.Borrower
		json.NewDecoder(resp.Body).Decode(&responseBorrower)
		assert.Equal(t, newBorrower.Name, responseBorrower.Name)
		assert.Equal(t, newBorrower.Email, responseBorrower.Email)
		assert.Equal(t, newBorrower.Phone, responseBorrower.Phone)
		newBorrower.Id = responseBorrower.Id
	})

	borrowerId := strconv.Itoa(newBorrower.Id)
	t.Run("Get Borrower by ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/borrowers/"+borrowerId, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseBorrower *domain.Borrower
		json.NewDecoder(resp.Body).Decode(&responseBorrower)
		assert.Equal(t, newBorrower, responseBorrower)
	})

	t.Run("UnregisterBorrower Borrower", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/borrowers/"+borrowerId, nil)
		resp, _ := app.Test(req)

		// Assert
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
