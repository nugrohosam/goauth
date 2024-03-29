package rolepermission

import (
	"github.com/nugrohosam/gosampleapi/pubsub/producers"
	"github.com/nugrohosam/gosampleapi/repositories/permission"
	"github.com/nugrohosam/gosampleapi/repositories/role"
	"gorm.io/gorm"
)

// TableName ..
const TableName = "role_permissions"

// RolePermission struct
type RolePermission struct {
	gorm.Model
	ID           int `gorm:"primaryKey;autoIncrement:true"`
	RoleID       int
	Role         role.Role `gorm:"constraint:OnDelete:CASCADE;references:id"`
	PermissionID int
	Permission   permission.Permission `gorm:"constraint:OnDelete:CASCADE;references:id"`
}

// RolePermissions using for many rolePermission_permissions
type RolePermissions []RolePermission

// RoleIDs ...
func RoleIDs(database *gorm.DB) *gorm.DB {
	query := database.Table(role.TableName)
	return query.Select("id")
}

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
	producers.ChangedRolePermission(rolePermission.ID)
	return nil
}
