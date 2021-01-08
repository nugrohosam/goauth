package v1

// CreateRolePermission ...
type CreateRolePermission struct {
	RoleID       string `json:"role_id" validate:"required,ascii"`
	PermisisonID string `json:"permission_id" validate:"required,ascii"`
}

// UpdateRolePermission ...
type UpdateRolePermission struct {
	RoleID       string `json:"role_id" validate:"required,ascii"`
	PermisisonID string `json:"permission_id" validate:"required,ascii"`
}
