\c postgres;

DROP DATABASE IF EXISTS library;
CREATE DATABASE library;

\c library;

CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    isbn VARCHAR(20) NOT NULL UNIQUE,
    stock INT NOT NULL
);

CREATE TABLE borrower (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20)
);

CREATE TABLE loan (
    id SERIAL PRIMARY KEY,
    book_isbn VARCHAR(20) NOT NULL,
    borrower_id INT NOT NULL,
    loan_date DATE NOT NULL,
    due_date DATE NOT NULL,
    return_date DATE,
    status VARCHAR(20) DEFAULT 'NOT_RETURNED',
    FOREIGN KEY (book_isbn) REFERENCES book(isbn),
    FOREIGN KEY (borrower_id) REFERENCES borrower(id)
);