package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword - function for hash password
func HashPassword(password string) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
  return string(bytes), err
}

// CheckPassword - function for verify password
func CheckPassword(hashedPassword, password string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
  return err == nil 
}
