package user

import (
	"errors"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Create using for user
func Create(name, username, email, password string) (User, error) {
	database := *conn.Db

	user := User{Name: name, Username: username, Email: email, Password: password}
	userExisting := User{}
	isExists := database.Where("email = ? OR username = ?", user.Email, user.Username).Find(&user).RowsAffected

	if isExists != 0 {
		return userExisting, errors.New("User is exists")
	}

	database.Create(&user)
	return user, nil
}

// Find is using
func Find(id string) User {
	database := *conn.Db

	user := User{}
	database.Where("id = ?", id).First(&user)

	return user
}

// FindByEmailOrUsername ...
func FindByEmailOrUsername(emailOrUsername string) User {
	database := *conn.Db

	user := User{}
	database.Where("username = ? OR email = ?", emailOrUsername, emailOrUsername).First(&user)

	return user
}
