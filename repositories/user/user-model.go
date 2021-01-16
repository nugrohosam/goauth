package user

// TableName ...
const TableName = "users"

// User struct
type User struct {
	ID          int
	Name        string
	Username    string
	Email       string
	Password    string
	Phonenumber string
}

// Users using for many users
type Users []User
