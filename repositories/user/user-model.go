package user

// TableName ...
const TableName = "users"

type UserID int

// User struct
type User struct {
	ID          UserID
	Name        string
	Username    string
	Email       string
	Password    string
	Phonenumber string
}

// Users using for many users
type Users []User
