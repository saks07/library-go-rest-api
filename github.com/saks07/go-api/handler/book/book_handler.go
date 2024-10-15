package bookhandler

import (
	"encoding/json"
	"net/http"
	"github.com/saks07/go-api/service/book"
   "github.com/saks07/go-api/utils"
)

type BookHandler struct {
   BookService bookservice.BookService
}

type Book struct {
   ID int `json:"id"`
   BookTitle string `json:"book_title"`
   BookAvailableCopies int `json:"book_available_copies"`
 }

func (h *BookHandler) ListBooksHandler(res http.ResponseWriter, req *http.Request) {
   if checkGet := utils.CheckGetMethod(req.Method); checkGet == false {
      http.Error(res, "Invalid request method", http.StatusBadRequest)
      return
   }

   books, err := h.BookService.ListBooks()

   if err != nil {
     http.Error(res, "Failed to list books", http.StatusInternalServerError)
     return
   }

   res.Header().Set("Content-Type", "application/json")
   json.NewEncoder(res).Encode(books)
}