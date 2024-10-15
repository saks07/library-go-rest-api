package userservice

import (
	"github.com/saks07/go-api/store/user"
)

type UserService struct {
	UserStore userstore.UserStore
}

func (s *UserService) CreateUser(username string, email string) error {
	return s.UserStore.SaveUser(username, email)
}

func (s *UserService) ListUsers() ([]userstore.User, error) {
	return s.UserStore.GetAllUsers()
}