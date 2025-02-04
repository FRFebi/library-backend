package http

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/FRFebi/library-backend/domain"
	"github.com/FRFebi/library-backend/infrastructure"
	"github.com/FRFebi/library-backend/repository"
	"github.com/FRFebi/library-backend/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func BookSetup(db *sql.DB) domain.BookUsecase {

	Bookrepo := repository.NewBookRepository(db)
	return usecase.NewBookUsecase(Bookrepo)
}

func TestBookHandler(t *testing.T) {
	app := fiber.New()
	api := app.Group("/api")
	db := infrastructure.NewPostgreeDB()
	bookUC := BookSetup(db)
	NewBookHandler(api, bookUC)

	random := strconv.Itoa(int(time.Now().Unix()))
	newBook := &domain.Book{Isbn: random, Title: "Unit Test", Author: "Tester", Stock: 2}
	t.Run("Create Book", func(t *testing.T) {
		bookJSON, _ := json.Marshal(newBook)

		req := httptest.NewRequest(http.MethodPost, "/api/books", bytes.NewBuffer(bookJSON))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseBook *domain.Book
		json.NewDecoder(resp.Body).Decode(&responseBook)
		assert.Equal(t, newBook.Title, responseBook.Title)
		assert.Equal(t, newBook.Author, responseBook.Author)
		assert.Equal(t, newBook.Isbn, responseBook.Isbn)
		newBook.Id = responseBook.Id
	})

	t.Run("Get Book by ISBN", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/books/"+newBook.Isbn, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var responseBook *domain.Book
		json.NewDecoder(resp.Body).Decode(&responseBook)
		assert.Equal(t, newBook, responseBook)
	})

	t.Run("Delete Book", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/books/"+newBook.Isbn, nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
