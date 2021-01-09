package v1

// Paginate ...
type Paginate struct {
	Items       interface{} `structs:"items" json:"items"`
	Total       int         `structs:"total" json:"total"`
	PerPage     int         `structs:"per_page" json:"per_page"`
	CurrentPage int         `structs:"current_page" json:"current_page"`
}
