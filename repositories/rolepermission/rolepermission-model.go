package rolepermission

import (
	"fmt"

	"github.com/nugrohosam/gosampleapi/listeners/producers"
	"github.com/nugrohosam/gosampleapi/repositories/permission"
	"github.com/nugrohosam/gosampleapi/repositories/role"
	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
	"gorm.io/gorm"
)

// TableName ..
const TableName = "role_permission"

// RolePermission struct
type RolePermission struct {
	ID           int
	RoleID       role.RoleID
	Role         role.Role `gorm:"constraint:OnDelete:CASCADE;references:id"`
	PermissionID permission.PermissionID
	Permission   permission.Permission `gorm:"constraint:OnDelete:CASCADE;references:id"`
}

// RolePermissions using for many rolePermission_permissions
type RolePermissions []RolePermission

// RoleIDs ...
func RoleIDs(database *gorm.DB) *gorm.DB {
	query := database.Table(role.TableName)
	return query.Select("id")
}

// PermissionIDs ...
func PermissionIDs(database *gorm.DB) *gorm.DB {
	query := database.Table(permission.TableName)
	return query.Select("id")
}

// RoleIDsUser ...
func RoleIDsUser(database *gorm.DB, userID string) *gorm.DB {
	return database.Table(userRoleRepo.TableName).Select("role_id").Where("user_id = ?", userID)
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
	producers.ChangedRolePermission(fmt.Sprint(rolePermission.ID))
	return nil
}
