package user

import (
	"errors"
	"strconv"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Get using for permission
func Get(search, limit, offset, orderBy string) (Users, int, error) {
	var permissions = Users{}
	database := *conn.DbOrm

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	totalRows := database.Where("name LIKE ? or email LIKE ?", "%"+search+"%", "%"+search+"%").Find(&permissions).RowsAffected
	database.Where("name LIKE ? or email LIKE ?", "%"+search+"%", "%"+search+"%").Limit(limitInt).Offset(offsetInt).Order("name " + orderBy).Find(&permissions)

	return permissions, int(totalRows), nil
}

// Create using for user
func Create(name, username, email, password string) (User, error) {
	database := *conn.DbOrm

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
func Find(ID int) User {
	database := *conn.DbOrm

	user := User{}
	database.Where("id = ?", ID).First(&user)

	return user
}

// FindByEmailOrUsername ...
func FindByEmailOrUsername(emailOrUsername string) User {
	database := *conn.DbOrm

	user := User{}
	database.Where("username = ? OR email = ?", emailOrUsername, emailOrUsername).First(&user)

	return user
}
