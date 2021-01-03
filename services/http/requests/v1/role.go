package v1

// CreateRole ...
type CreateRole struct {
	Name string `json:"name" validate:"required,ascii"`
}

// UpdateRole ...
type UpdateRole struct {
	Name string `json:"name" validate:"required,ascii"`
}
