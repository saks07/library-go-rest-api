package borrowedhandler

import (
	"encoding/json"
	"net/http"
	"github.com/saks07/go-api/service/borrowed_books"
	"github.com/saks07/go-api/utils"
)

type BorrowedHandler struct {
   BorrowedService borrowedservice.BorrowedService
}

type BorrowedBook struct {
  ID int `json:"id"`
  UserID int `json:"user_id"`
  BookID int `json:"book_id"`
	BorrowDate int `json:"borrow_date"`
	ReturnDate int `json:"return_date"`
}

func (h *BorrowedHandler) ListBorrowedBooksHandler(res http.ResponseWriter, req *http.Request) {
	if checkGet := utils.CheckGetMethod(req.Method); checkGet == false {
    http.Error(res, "Invalid request method", http.StatusBadRequest)
    return
  }

	var userId string = req.PathValue("userId");

	// Validate URI param user ID
	if err := utils.IsNumber(userId); err == false {
		http.Error(res, "Invalid user ID", http.StatusBadRequest)
		return
	}

	borrowed, err := h.BorrowedService.ListBorrowedBooks(userId)

	if err != nil {
		http.Error(res, "Failed to list borrowed books", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(borrowed)
}

func (h *BorrowedHandler) ListReturnedBooksHandler(res http.ResponseWriter, req *http.Request) {
	if checkGet := utils.CheckGetMethod(req.Method); checkGet == false {
    http.Error(res, "Invalid request method", http.StatusBadRequest)
    return
  }

	var userId string = req.PathValue("userId");

	// Validate URI param user ID
	if err := utils.IsNumber(userId); err == false {
		http.Error(res, "Invalid user ID", http.StatusBadRequest)
		return
	}

	borrowed, err := h.BorrowedService.ListReturnedBooks(userId)

	if err != nil {
		http.Error(res, "Failed to list returned books", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(borrowed)
}