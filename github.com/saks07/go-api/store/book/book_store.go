package bookstore

import (
	"database/sql"
	_ "github.com/lib/pq"
  "github.com/saks07/go-api/utils"
)

type Book struct {
	ID int `json:"id"`
  BookTitle string `json:"book_title"`
  BookAvailableCopies int `json:"book_available_copies"`
 }

type BookStore interface {
   GetAllBooks() ([]Book, error)
}

type SQLBookStore struct {
   DB *sql.DB
}

// Global variables
var dbTable string = "books"

func (s *SQLBookStore) GetAllBooks() ([]Book, error) {
  var query string = utils.QueryStringTable("SELECT id, book_title, book_available_copies FROM {table}", dbTable)
  rows, err := s.DB.Query(query)
  if err != nil {
    return nil, err
  }

  defer rows.Close()

  var books []Book

  for rows.Next() {
    var book Book

    if err := rows.Scan(&book.ID, &book.BookTitle, &book.BookAvailableCopies); err != nil {
      return nil, err
    }

    books = append(books, book)
  }

  return books, nil
}