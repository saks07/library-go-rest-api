package borrowedstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/saks07/go-api/utils"
	"time"
	"fmt"
)

type BorrowedBook struct {
  ID int `json:"id"`
  UserID int `json:"user_id"`
  BookID int `json:"book_id"`
	BorrowDate string `json:"borrow_date"`
	ReturnDate *string `json:"return_date"` // NULL - if book is not yet returned
	Username string `json:"username"`
	Email string `json:"email"`
	BookTitle string `json:"book_title"`
	BookAvailableCopies int `json:"book_available_copies"`
}

type BorrowedStore interface {
	GetUserBorrowedBooks(userId string) ([]BorrowedBook, error)
	GetUserReturnedBooks(userId string) ([]BorrowedBook, error)
	SaveBorrowedBooks(bookId int, userId int) error
	UpdateReturnedBooks(id int) error
}

type SQLBorrowedStore struct {
   DB *sql.DB
}

// Global variables
var dbTable string = "borrowed_books"

func (s *SQLBorrowedStore) GetUserBorrowedBooks(userId string) ([]BorrowedBook, error) {

	var query string = utils.QueryStringTable("SELECT borrowed_books.id, borrowed_books.book_id, borrowed_books.user_id, borrowed_books.borrow_date, borrowed_books.return_date, users.username, users.email, books.book_title, books.book_available_copies FROM {table} JOIN users ON borrowed_books.user_id=users.id JOIN books ON borrowed_books.book_id=books.id WHERE borrowed_books.user_id=$1", dbTable);

	rows, err := s.DB.Query(query, userId)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

   defer rows.Close()

   var borrowedBooks []BorrowedBook

   for rows.Next() {
     var borrowedBook BorrowedBook

     if err := rows.Scan(&borrowedBook.ID, &borrowedBook.BookID, &borrowedBook.UserID, &borrowedBook.BorrowDate, &borrowedBook.ReturnDate, &borrowedBook.Username, &borrowedBook.Email, &borrowedBook.BookTitle, &borrowedBook.BookAvailableCopies); err != nil {
       return nil, err
     }

     borrowedBooks = append(borrowedBooks, borrowedBook)
   }

   return borrowedBooks, nil
}

func (s *SQLBorrowedStore) GetUserReturnedBooks(userId string) ([]BorrowedBook, error) {

	var query string = utils.QueryStringTable("SELECT borrowed_books.id, borrowed_books.book_id, borrowed_books.user_id, borrowed_books.borrow_date, borrowed_books.return_date, users.username, users.email, books.book_title, books.book_available_copies FROM {table} JOIN users ON borrowed_books.user_id=users.id JOIN books ON borrowed_books.book_id=books.id WHERE borrowed_books.user_id=$1 AND borrowed_books.return_date IS NOT NULL", dbTable);

	rows, err := s.DB.Query(query, userId)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

   defer rows.Close()

   var borrowedBooks []BorrowedBook

   for rows.Next() {
     var borrowedBook BorrowedBook

     if err := rows.Scan(&borrowedBook.ID, &borrowedBook.BookID, &borrowedBook.UserID, &borrowedBook.BorrowDate, &borrowedBook.ReturnDate, &borrowedBook.Username, &borrowedBook.Email, &borrowedBook.BookTitle, &borrowedBook.BookAvailableCopies); err != nil {
       return nil, err
     }

     borrowedBooks = append(borrowedBooks, borrowedBook)
   }

   return borrowedBooks, nil
}

func (s *SQLBorrowedStore) SaveBorrowedBooks(bookId int, userId int) error {
	var query string = utils.QueryStringTable("INSERT INTO {table} (book_id, user_id, borrow_date, return_date) VALUES ($1, $2, $3, NULL)", dbTable)
	stmt, stmtErr := s.DB.Prepare(query)
	
	if stmtErr != nil {
	 return stmtErr
 }

	defer stmt.Close()

	formatted := createDateFormatted();
	_, err := stmt.Exec(bookId, userId, formatted)

	return err
}

func (s *SQLBorrowedStore) UpdateReturnedBooks(id int) error {
	var query string = utils.QueryStringTable("UPDATE {table} SET return_date = $1 WHERE id = $2", dbTable)
	stmt, stmtErr := s.DB.Prepare(query)
	
	if stmtErr != nil {
	 return stmtErr
 }

	defer stmt.Close()

	formatted := createDateFormatted();
	_, err := stmt.Exec(formatted, id)

	return err
}

func createDateFormatted() string {
	timeNow := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", timeNow.Year(), timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second())
}