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

// UsetListItems ..
type UsetListItems []UserItems

// UserPaginate ...
type UserPaginate struct {
	Items       UsetListItems `structs:"items" json:"items"`
	Total       int           `structs:"total" json:"total"`
	PerPage     int           `structs:"per_page" json:"per_page"`
	CurrentPage int           `structs:"current_page" json:"current_page"`
}
