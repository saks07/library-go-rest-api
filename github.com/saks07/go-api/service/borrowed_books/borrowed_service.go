package borrowedservice

import (
	"github.com/saks07/go-api/store/borrowed_books"
)

type BorrowedService struct {
	BorrowedStore borrowedstore.BorrowedStore
}

func (s *BorrowedService) ListBorrowedBooks(userId string) ([]borrowedstore.BorrowedBook, error) {
	return s.BorrowedStore.GetUserBorrowedBooks(userId)
}

func (s *BorrowedService) ListReturnedBooks(userId string) ([]borrowedstore.BorrowedBook, error) {
	return s.BorrowedStore.GetUserReturnedBooks(userId)
}