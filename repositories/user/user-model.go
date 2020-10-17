package user

// TableName ...
const TableName = "users"

// User struct
type User struct {
	ID       int32
	Name     string
	Username string
	Email    string
	Password string
}

// Users using for many users
type Users []User
