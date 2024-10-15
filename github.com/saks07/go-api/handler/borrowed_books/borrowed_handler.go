package borrowedhandler

import (
	"encoding/json"
	"net/http"
	"github.com/saks07/go-api/service/borrowed_books"
	"github.com/saks07/go-api/utils"
	"time"
)

type BorrowedHandler struct {
   BorrowedService borrowedservice.BorrowedService
}

type BorrowedBook struct {
  ID int `json:"id"`
  UserID int `json:"user_id"`
  BookID int `json:"book_id"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date"`
}

type BorrowedBookInsert struct {
  UserID int `json:"user_id"`
  BookID int `json:"book_id"`
}

type BorrowedBookUpdate struct {
  ID int `json:"id"`
}

func (h *BorrowedHandler) ListBorrowedBooksHandler(res http.ResponseWriter, req *http.Request) {
	if checkGet := utils.CheckGetMethod(req.Method); checkGet == false {
    http.Error(res, "Invalid request method", http.StatusBadRequest)
    return
  }

	var userId string = req.PathValue("userId");

	// Validate URI param user ID
	if err := utils.IsStringNumber(userId); err == false {
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
	if err := utils.IsStringNumber(userId); err == false {
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

func (h *BorrowedHandler) CreateBorrowedBooksHandler(res http.ResponseWriter, req *http.Request) {
  if checkPost := utils.CheckPostMethod(req.Method); checkPost == false {
    http.Error(res, "Invalid request method", http.StatusBadRequest)
    return
  }

  var borrowedBook BorrowedBookInsert

  if err := json.NewDecoder(req.Body).Decode(&borrowedBook); err != nil {
    http.Error(res, "Invalid request body", http.StatusBadRequest)
    return
  }

  if err := h.BorrowedService.CreateBorrowedBooks(borrowedBook.BookID, borrowedBook.UserID); err != nil {
    http.Error(res, "Failed to create borrowed book", http.StatusInternalServerError)
    return
  }

  res.WriteHeader(http.StatusCreated)
}

func (h *BorrowedHandler) UpdateReturnedBooksHandler(res http.ResponseWriter, req *http.Request) {
  if checkPut := utils.CheckPutMethod(req.Method); checkPut == false {
    http.Error(res, "Invalid request method", http.StatusBadRequest)
    return
  }

	var returnedBook BorrowedBookUpdate

  if err := json.NewDecoder(req.Body).Decode(&returnedBook); err != nil {
    http.Error(res, "Invalid request body", http.StatusBadRequest)
    return
  }

  if err := h.BorrowedService.UpdateBorrowedBooks(returnedBook.ID); err != nil {
    http.Error(res, "Failed to update borrowed book", http.StatusInternalServerError)
    return
  }

  res.WriteHeader(http.StatusCreated)
}