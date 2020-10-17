package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// MakeHash ...
func MakeHash(data string) string {
	byted := []byte(data)
	hash, err := bcrypt.GenerateFromPassword(byted, bcrypt.MinCost)
	if err != nil {
		panic("Cannot hash")
	}

	return string(hash)
}

// CompareHash ...
func CompareHash(hashedPassword []byte, plainPassword string) bool {

	byted2 := []byte(plainPassword)

	err := bcrypt.CompareHashAndPassword(hashedPassword, byted2)
	if err != nil {
		return false
	}

	return true
}
