package main

import (
	"log"
	"os"

	"github.com/FRFebi/library-backend/delivery/http"
	"github.com/FRFebi/library-backend/infrastructure"
	"github.com/FRFebi/library-backend/repository"
	"github.com/FRFebi/library-backend/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	db := infrastructure.NewPostgreeDB()

	// Dependency Injection
	//1. Book
	bookRepo := repository.NewBookRepository(db)
	bookUC := usecase.NewBookUsecase(bookRepo)
	http.NewBookhandler(api, bookUC)

	//2. Borrower
	borrowerRepo := repository.NewBorrowerRepository(db)
	borrowerUC := usecase.NewBorrowerUsecase(borrowerRepo)
	http.NewBorrowerHandler(api, borrowerUC)

	//3. Loan
	loanRepo := repository.NewLoanRepository(db)
	loanUC := usecase.NewLoanUsecase(loanRepo)
	http.NewLoanHandler(api, loanUC)

	log.Fatal(app.Listen(":" + port))

}
