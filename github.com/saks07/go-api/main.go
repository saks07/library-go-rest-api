package main

import (
	"database/sql"
  "log"
	"net/http"
	"github.com/saks07/go-api/handler/user"
	"github.com/saks07/go-api/service/user"
	"github.com/saks07/go-api/store/user"
	"github.com/joho/godotenv"
	"os"
	"github.com/saks07/go-api/handler/book"
	"github.com/saks07/go-api/service/book"
	"github.com/saks07/go-api/store/book"
	"github.com/saks07/go-api/handler/borrowed_books"
	"github.com/saks07/go-api/service/borrowed_books"
	"github.com/saks07/go-api/store/borrowed_books"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var PORT string = ":8080"

func main() {
	//load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB credentials
	var dbHost string = os.Getenv("DB_HOST")
	var dbPort string = os.Getenv("DB_PORT")
	var dbUsername string = os.Getenv("DB_USERNAME")
	var dbPassword string = os.Getenv("DB_PASSWORD")
	var dbName string = os.Getenv("DB_NAME")

	// Database connection
	dataSourceName := "postgres://" + dbUsername + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not start SQL driver: %v", err)
	}
 
	m, err := migrate.NewWithDatabaseInstance("file://database/migrations","postgres", driver)
	if err != nil {
		log.Fatalf("Could not start migration: %v", err)
	}
 
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	// Initialize user store
	userStore := &userstore.SQLUserStore{DB: db}
	userService := userservice.UserService{UserStore: userStore}
	userHandler := userhandler.UserHandler{UserService: userService}

	// User API endpoints
	http.HandleFunc("/users/add", userHandler.CreateUserHandler)
  http.HandleFunc("/users/list", userHandler.ListUsersHandler)

	// Initialize book store
	bookStore := &bookstore.SQLBookStore{DB: db}
	bookService := bookservice.BookService{BookStore: bookStore}
	bookHandler := bookhandler.BookHandler{BookService: bookService}

	// Book API endpoints
	http.HandleFunc("/books/list", bookHandler.ListBooksHandler)

	// Initialize borrowed books store
	borrowedStore := &borrowedstore.SQLBorrowedStore{DB: db}
	borrowedService := borrowedservice.BorrowedService{BorrowedStore: borrowedStore}
	borrowedHandler := borrowedhandler.BorrowedHandler{BorrowedService: borrowedService}

	// Borrowed books API endpoints
	http.HandleFunc("/borrowed/{userId}", borrowedHandler.ListBorrowedBooksHandler)
	http.HandleFunc("/returned/{userId}", borrowedHandler.ListReturnedBooksHandler)
  
	// Start server
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}