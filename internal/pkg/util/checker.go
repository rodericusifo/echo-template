package util

import (
	"golang.org/x/crypto/bcrypt"
)

func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
