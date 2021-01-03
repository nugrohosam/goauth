package rolepermission

import (
	"github.com/nugrohosam/gosampleapi/repositories/permission"
	"github.com/nugrohosam/gosampleapi/repositories/role"
)

// TableName ...
const TableName = "role_permission"

// RolePermission struct
type RolePermission struct {
	ID           int
	RoleID       int
	Role         role.Role `gorm:"constraint:OnDelete:CASCADE;references:id"`
	PermissionID int
	Permission   permission.Permission `gorm:"constraint:OnDelete:CASCADE;references:id"`
}

// RolePermissions using for many role_permissions
type RolePermissions []RolePermission
