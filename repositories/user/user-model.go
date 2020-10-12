package user

// User struct
type User struct {
	ID   int32
	Name string
}

// Users using for many users
type Users []User
