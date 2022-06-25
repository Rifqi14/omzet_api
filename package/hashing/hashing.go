package hashing

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	return string(hash), err
}

func CheckHashString(password, hash string) bool {
	hasher := md5.New()
	hasher.Write([]byte(password))
	passwordHash := fmt.Sprintf("%x", hasher.Sum(nil))
	if condition := passwordHash == hash; condition {
		return true
	}
	return false
}
