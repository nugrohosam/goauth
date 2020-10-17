package usecases

import (
	"time"
	"errors"
	viper "github.com/spf13/viper"

	jwt "github.com/dgrijalva/jwt-go"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
)

// AuthBasic ...
func AuthBasic(emailOrUsername, password string) (string, error) {
	hashedPassword := helpers.MakeHash(password)
	user := userRepo.FindByEmailOrUsernameAndPassword(emailOrUsername, hashedPassword)

	if len(user.Username) == 0 || len(user.Email) == 0 {
		return "", errors.New("Cannot find user")
	}

	if isPasswordValid := helpers.CompareHash([]byte(user.Password), password); isPasswordValid {
		return "", errors.New("Cannot find user")
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email": user.Email,
		"expiredTime": time.Now().AddDate(1, 0, 0),
	})

	secret := viper.GetString("secret")
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", errors.New("Cannot make token")
	}

	return tokenString, nil
}
