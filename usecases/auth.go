package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/viper"

	jwt "github.com/dgrijalva/jwt-go"
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
)

// AuthBasic ...
func AuthBasic(emailOrUsername, password string) (string, error) {
	user := userRepo.FindByEmailOrUsername(emailOrUsername)

	if len(user.Username) == 0 || len(user.Email) == 0 {
		return "", errors.New("cannot find user, username or email")
	}

	if isPasswordValid := helpers.CompareHash(user.Password, password); !isPasswordValid {
		return "", errors.New("cannot find user, password")
	}

	tokenExpiredInHour, _ := strconv.ParseInt(viper.GetString("token.expired_time"), 24, 64)
	dataUser := map[string]interface{}{
		"id":       user.ID,
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
	}

	dataTokenString, _ := json.Marshal(dataUser)
	secret := viper.GetString("secret")
	dataTokenEncrypted := helpers.Encrypt(string(dataTokenString), secret)

	dataToken := map[string]interface{}{
		"data":        dataTokenEncrypted,
		"expiredTime": time.Now().Add(time.Hour * time.Duration(tokenExpiredInHour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(dataToken))

	bytedString := helpers.GetBytedSecret()
	tokenString, err := token.SignedString(bytedString) // always use byted string

	if err != nil {
		return "", errors.New("cannot make token")
	}

	return tokenString, nil
}

// AuthorizationValidation ...
func AuthorizationValidation(tokenString string) error {
	token, err := jwt.Parse(tokenString, validateToken)
	if err != nil {
		return errors.New("wrong token input")
	}

	if data, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		timeNow := time.Now()
		timeExpiredString := data["expiredTime"].(string)
		layout := "2006-01-02T15:04:05+07:00"
		timeExpired, _ := time.Parse(layout, timeExpiredString)

		if timeNow.Before(timeExpired) {
			return nil
		}

		return errors.New("time token has expired")
	}

	return errors.New("cannot validate")
}

// GetDataAuth ...
func GetDataAuth(tokenString string) (map[string]interface{}, error) {

	token, err := jwt.Parse(tokenString, validateToken)
	dataToken, ok := token.Claims.(jwt.MapClaims)

	secret := viper.GetString("secret")
	dataUserInString := helpers.Decrypt(dataToken["data"].(string), secret)

	var dataUser map[string]interface{}
	json.Unmarshal([]byte(dataUserInString), &dataUser)

	if ok && token.Valid {
		return dataUser, nil
	} else if err != nil {
		return nil, errors.New("cannot validate auth token")
	} else {
		return nil, errors.New("cannot validate")
	}
}

func validateToken(token *jwt.Token) (interface{}, error) {
	bytedString := helpers.GetBytedSecret()

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}

	return bytedString, nil
}
