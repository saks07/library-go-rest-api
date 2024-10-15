package userhandler

import (
	"encoding/json"
	"net/http"
	"github.com/saks07/go-api/service/user"
  "github.com/saks07/go-api/utils"
)

type UserHandler struct {
   UserService userservice.UserService
}

type User struct {
  ID int `json:"id"`
  Username string `json:"username"`
  Email string `json:"email"`
}

func (h *UserHandler) CreateUserHandler(res http.ResponseWriter, req *http.Request) {
  if checkPost := utils.CheckPostMethod(req.Method); checkPost == false {
    http.Error(res, "Invalid request method", http.StatusBadRequest)
    return
  }

  var user User

  if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
    http.Error(res, "Invalid request body", http.StatusBadRequest)
    return
  }

  if user.Username == "" || user.Email == "" {
    http.Error(res, "Missing username or email", http.StatusBadRequest)
    return
  }

  if err := utils.ValidEmail(user.Email); err == false {
    http.Error(res, "Invalid email address", http.StatusBadRequest)
    return
  }

  if err := h.UserService.CreateUser(user.Username, user.Email); err != nil {
    http.Error(res, "Failed to create user", http.StatusInternalServerError)
    return
  }

  res.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) ListUsersHandler(res http.ResponseWriter, req *http.Request) {
  if checkGet := utils.CheckGetMethod(req.Method); checkGet == false {
    http.Error(res, "Invalid request method", http.StatusBadRequest)
    return
  }

  users, err := h.UserService.ListUsers()

  if err != nil {
    http.Error(res, "Failed to list users", http.StatusInternalServerError)
    return
  }

  res.Header().Set("Content-Type", "application/json")
  json.NewEncoder(res).Encode(users)
}