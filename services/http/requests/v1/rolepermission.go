package v1

// CreateRolePermission ...
type CreateRolePermission struct {
	RoleID       interface{} `json:"role_id" validate:"required,min=1,should-be-integer"`
	PermisisonID interface{} `json:"permission_id" validate:"required,min=1,should-be-integer"`
}
