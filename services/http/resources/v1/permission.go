package v1

// PermissionItems ...
type PermissionItems struct {
	Name string `copier:"must" structs:"name" json:"name" deepcopier:"field:Name"`
}

// PermissionDetail ...
type PermissionDetail struct {
	ID   string `structs:"id" json:"id" copier:"field:ID"`
	Name string `structs:"name" json:"name" copier:"field:Name"`
}

// PermissionListItems ..
type PermissionListItems []PermissionItems

// PermissionPaginate ...
type PermissionPaginate struct {
	Items       PermissionListItems `structs:"items" json:"item"`
	Total       int                 `structs:"total" json:"total"`
	PerPage     int                 `structs:"per_page" json:"per_page"`
	CurrentPage int                 `structs:"current_page" json:"current_page"`
}
