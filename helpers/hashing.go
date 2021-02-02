package helpers

import (
	viper "github.com/spf13/viper"
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
	bytedPlainPassword := []byte(plainPassword)
	err := bcrypt.CompareHashAndPassword(bytedPlainPassword, hashedPassword)
	if err != nil {
		return false
	}

	return true
}

// GetBytedSecret ...
func GetBytedSecret() []byte {
	secret := viper.GetString("secret")
	return []byte(secret)
}
