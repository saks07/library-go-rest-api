package utils

import (
	"net/mail"

	"strconv"
)
	

func ValidEmail(email string) bool {
  _, err := mail.ParseAddress(email)
  return err == nil
}

func IsStringNumber(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}