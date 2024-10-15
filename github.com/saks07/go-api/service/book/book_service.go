package bookservice

import (
	"github.com/saks07/go-api/store/book"
)

type BookService struct {
	BookStore bookstore.BookStore
}

func (s *BookService) ListBooks() ([]bookstore.Book, error) {
	return s.BookStore.GetAllBooks()
}