package rolepermission

import (
	"fmt"

	"github.com/nugrohosam/gosampleapi/listeners/producers"
	"github.com/nugrohosam/gosampleapi/repositories/permission"
	"github.com/nugrohosam/gosampleapi/repositories/role"
	"gorm.io/gorm"
)

// TableName ..
const TableName = "role_permission"

// RolePermission struct
type RolePermission struct {
	ID           int
	RoleID       int
	Role         role.Role `gorm:"constraint:OnDelete:CASCADE;references:id"`
	PermissionID int
	Permission   permission.Permission `gorm:"constraint:OnDelete:CASCADE;references:id"`
}

// RolePermissions using for many rolePermission_permissions
type RolePermissions []RolePermission

// AfterCreate ..
func (rolePermission *RolePermission) AfterCreate(tx *gorm.DB) error {
	return nil
}

// AfterDelete ..
func (rolePermission *RolePermission) AfterDelete(tx *gorm.DB) error {
	return nil
}

// AfterUpdate ..
func (rolePermission *RolePermission) AfterUpdate(tx *gorm.DB) error {
	producers.ChangedRolePermission(fmt.Sprint(rolePermission.ID))
	return nil
}
