package user

import "gorm.io/gorm"

// TableName ...
const TableName = "users"

// User struct
type User struct {
	gorm.Model
	ID          int `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Username    string
	Email       string
	Password    string
	Phonenumber string
}

// Users using for many users
type Users []User
