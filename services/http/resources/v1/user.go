package v1

// UserItems ...
type UserItems struct {
	Name string `structs:"name" json:"name"`
}

// UserDetail ...
type UserDetail struct {
	Name  string `structs:"name" json:"name"`
	Email string `structs:"email" json:"email"`
}

// UserListItems ..
type UserListItems []UserItems
