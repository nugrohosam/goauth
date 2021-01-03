package userrole

import (
	"github.com/nugrohosam/gosampleapi/repositories/role"
	"github.com/nugrohosam/gosampleapi/repositories/user"
)

// TableName ...
const TableName = "user_role"

// UserRole struct
type UserRole struct {
	ID     int
	UserID int
	User   user.User `gorm:"constraint:OnDelete:CASCADE;references:id"`
	RoleID int
	Role   role.Role `gorm:"constraint:OnDelete:CASCADE;references:id"`
}

// RolePermissions using for many role_permissions
type RolePermissions []UserRole
