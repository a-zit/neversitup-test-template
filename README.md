# neversitup-test-template

## Project Structure

Concept of microservices & clean architecture with Golang

```
book-backend/
├── LICENSE
├── Makefile -- Makefile for common tasks
├── README.md
├── cmd
│   ├── backend-bo-api
│   │   ├── config
│   │   │   └── config.go -- Configuration file
│   │   ├── handler
│   │   │   └── book.go -- HTTP handler
│   │   ├── main.go
│   │   └── middleware
│   │       └── error.go -- Error handling middleware
│   └── backend-client-api
│       └── main.go
├── database -- Database configuration and scripts
│   ├── mariadb.go
│   └── script.sql
├── docker-compose.yml
├── domain
│   ├── book.go -- Book entity, use-case, repository interface
│   ├── common.go
│   ├── error.go
│   └── mocks
│       └── BookRepository.go -- Mock repository for testing
├── go.mod
├── go.sum
├── module
│   └── book
│       ├── book_suite
│       │   └── book_suite.go -- Test suite for book
│       ├── repository.go -- repository for book
│       ├── usecase.go -- business logic for book
│       └── usecase_test.go
├── pkg
│   ├── error
│   │   └── app_error.go -- Custom Application error
│   └── validator
│       └── validator.go -- Request validation
```
