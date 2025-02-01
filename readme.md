# Library Management System

## Description

This library management system is designed to assist library administrators in streamlining the book borrowing process. The available features include:

1. Input book data
2. Input borrower data
3. Borrowers can borrow and return books
4. Administrators can determine whether borrowed books are returned on time or late

## Additional Assumptions

1. Each book can only be borrowed by one borrower at a time.
2. The return status of a book is calculated based on the return date and the due date.
3. The return status of a book will automatically be updated to `ON_TIME` if the book is returned before or on the due date, and `LATE` if returned after the due date.
4. The return status of a book that has not been returned will remain `NOT_RETURNED`.
5. Borrower and book data must be valid before a loan is processed.
6. A borrowed book cannot be borrowed by another borrower until it is returned.

## How to Run the Application

1. Ensure that the MySQL server is installed and running.
2. Create a new database named `library_management_system`.
3. Adjust the database configuration in the `config/config.go` file according to your MySQL setup.
4. Run the following commands to install dependencies:
5.

```bash
 go mod tidy
 go run cmd/main.go
```

5. Access the API via the endpoint: `http://localhost:8080`

6. **API Endpoints**:

   1. **Book**
      | Method | Endpoint | Description | Body (JSON) |
      | ------ | ---------------- | ------------------------------------ | --------------------------------------------------------------------------- |
      | GET | `/api/books` | Get all books | - |
      | GET | `/api/books/{id}`| Get a book by ID | - |
      | POST | `/api/books` | Add a new book | `{ "title": "string", "author": "string", "isbn": "string", "available": true/false }` |
      | DELETE | `/api/books/{id}`| Delete a book by ID | - |

   2. **Borrower**
      | Method | Endpoint | Description | Body (JSON) |
      | ------ | ------------------ | ------------------------------------ | -------------------------------------------------------- |
      | GET | `/api/borrowers` | Get all borrowers | - |
      | GET | `/api/borrowers/{id}` | Get a borrower by ID | - |
      | POST | `/api/borrowers` | Add a new borrower | `{ "name": "string", "email": "string", "phone": "string" }` |
      | DELETE | `/api/borrowers/{id}` | Delete a borrower by ID | - |

   3. **Loan**
      | Method | Endpoint | Description | Body (JSON) |
      | ------ | ------------------ | ------------------------------------ | --------------------------------------------------------------------------- |
      | GET | `/api/loans` | Get all loans | - |
      | GET | `/api/loans/{id}` | Get a loan by ID | - |
      | POST | `/api/loans/borrow`| Borrow a book | `{ "bookID": int, "borrowerID": int, "loanDate": "YYYY-MM-DD", "dueDate": "YYYY-MM-DD" }` |
      | POST | `/api/loans/return`| Return a book | `{ "loanID": int, "returnDate": "YYYY-MM-DD" }` |

---

## Table Information

### `book` Table

| Column    | Data Type    | Description                 |
| --------- | ------------ | --------------------------- |
| id        | INT          | Primary Key, Auto Increment |
| title     | VARCHAR(255) | Book title                  |
| author    | VARCHAR(255) | Book author                 |
| isbn      | VARCHAR(20)  | Book ISBN, must be unique   |
| available | BOOLEAN      | Book availability status    |

### `borrower` Table

| Column | Data Type    | Description                    |
| ------ | ------------ | ------------------------------ |
| id     | INT          | Primary Key, Auto Increment    |
| name   | VARCHAR(255) | Borrower name                  |
| email  | VARCHAR(255) | Borrower email, must be unique |
| phone  | VARCHAR(20)  | Borrower phone number          |

### `loan` Table

| Column      | Data Type | Description                                     |
| ----------- | --------- | ----------------------------------------------- |
| id          | INT       | Primary Key, Auto Increment                     |
| book_id     | INT       | Foreign Key, Book ID                            |
| borrower_id | INT       | Foreign Key, Borrower ID                        |
| loan_date   | DATE      | Loan date                                       |
| due_date    | DATE      | Due date                                        |
| return_date | DATE      | Return date (nullable)                          |
| status      | ENUM      | Loan status (`ON_TIME`, `LATE`, `NOT_RETURNED`) |
