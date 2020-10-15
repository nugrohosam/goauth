package v1

// UserItems ...
type UserItems struct {
	Name string
}

// UserDetail ...
type UserDetail struct {
	Name  string
	Email string
}

// UserPaginate ...
type UserPaginate struct {
	Item        []UserItems
	Total       int
	PerPage     int
	CurrentPage int
}
