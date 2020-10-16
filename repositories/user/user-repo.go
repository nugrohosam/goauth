package user

import (
	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// GetAll using
func GetAll() {

}

// Create using for user
func Create(name, username, email, password string) User {
	database := *conn.Db

	user := User{Name: name, Username: username, Email: email, Password: password}
	database.Create(&user)

	return user
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
