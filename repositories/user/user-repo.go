package user

import (
	"errors"
	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// GetAll using
func GetAll() {

}

// Create using for user
func Create(name, username, email, password string) (User, error) {
	database := *conn.Db

	user := User{Name: name, Username: username, Email: email, Password: password}
	userExisting := User{}
	
	var isExists int64
	database.Where("email = ? or usernme = = ? ",  user.Email, user.Username).Count(&isExists)

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

// FindByEmailOrUsernameAndPassword ...
func FindByEmailOrUsernameAndPassword(emailOrUsername, password string) User {
	database := *conn.Db

	user := User{}
	database.Where("username = ? OR email = ?", emailOrUsername, emailOrUsername).Where("password = ?", password).First(&user)

	return user
}
