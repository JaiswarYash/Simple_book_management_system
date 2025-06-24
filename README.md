# Book Management System

This is a simple book management system built with Go.

## Features

- Create, Read, Update, and Delete books.
- API endpoints for managing books.

## Getting Started

### Prerequisites

- Go (version 1.15 or higher)
- Postman (or any other API client)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mr-yash-dev/Book-management-system.git
   ```
2. Navigate to the project directory:
   ```bash
   cd Book-management-system
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the Application

To run the application, execute the following command:

```bash
go run cmd/server/main.go
```

The server will start on `localhost:9010`.

## API Endpoints

The following endpoints are available:

- `GET /api/v1/books`: Get all books
- `GET /api/v1/books/{bookId}`: Get a book by its ID
- `POST /api/v1/books`: Create a new book
- `PUT /api/v1/books/{bookId}`: Update a book
- `DELETE /api/v1/books/{bookId}`: Delete a book 