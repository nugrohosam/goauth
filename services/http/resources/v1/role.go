package v1

// RoleItem ...
type RoleItem struct {
	ID   int    `structs:"id" json:"id" copier:"field:ID"`
	Name string `structs:"name" json:"name" copier:"field:Name"`
}

// RoleDetail ...
type RoleDetail struct {
	ID   int    `structs:"id" json:"id" copier:"field:ID"`
	Name string `structs:"name" json:"name" copier:"field:Name"`
}

// RoleListItems ..
type RoleListItems []RoleItem
