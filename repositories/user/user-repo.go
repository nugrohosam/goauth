package user

import (
	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// GetAll using
func GetAll() {

}

// Create using for user
func Create(name string) User {
	database := *conn.Db

	user := User{Name: name}
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
