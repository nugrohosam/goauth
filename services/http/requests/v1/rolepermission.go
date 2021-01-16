package v1

// CreateRolePermission ...
type CreateRolePermission struct {
	RoleID       int `form:"role_id" json:"role_id" binding:"required,numeric"`
	PermisisonID int `form:"permission_id" json:"permission_id" binding:"required,numeric"`
}
