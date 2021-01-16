package v1

// CreateUserRolePermission ...
type CreateUserRolePermission struct {
	UserID int `form:"user_id" json:"user_id" binding:"required,numeric"`
	RoleID int `form:"role_id" json:"role_id" binding:"required,numeric"`
}
