package v1

// UserItems ...
type UserItems struct {
	Name string `json:"name"`
}

// UserDetail ...
type UserDetail struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserPaginate ...
type UserPaginate struct {
	Item        []UserItems
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
}
