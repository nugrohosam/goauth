package rolepermission

import (
	"errors"
	"strconv"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
	"gorm.io/gorm/clause"
)

// Get using for rolePermission
func Get(search, limit, offset, orderBy string) (RolePermissions, int, error) {
	var rolePermissions = RolePermissions{}
	database := *conn.DbOrm

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	roleIds := RoleIDs(&database).Where("name LIKE ?", "%"+search+"%")
	totalRows := database.Table(TableName).Where("id IN (?)", roleIds).Find(&RolePermissions{}).RowsAffected
	database.Table(TableName).Where("role_id IN (?)", roleIds).Limit(limitInt).Offset(offsetInt).Order("id " + orderBy).Find(&rolePermissions)

	return rolePermissions, int(totalRows), nil
}

// GetByRoleIDs is using
func GetByRoleIDs(roleIDs []int) RolePermissions {
	database := *conn.DbOrm

	var rolePermissions = RolePermissions{}
	database.Table(TableName).Where("role_id IN ?", roleIDs).Find(&rolePermissions)

	return rolePermissions
}

// FindWithID is using
func FindWithID(ID string) RolePermission {
	database := *conn.DbOrm

	rolePermission := RolePermission{}
	database.Table(TableName).Preload(clause.Associations).Where("id = ?", ID).First(&rolePermission)

	return rolePermission
}

// FindByUserIDAndPermissionName is using
func FindByUserIDAndPermissionName(userID string, permissionName []string) RolePermission {
	database := *conn.DbOrm

	rolePermission := RolePermission{}

	roleIds := RoleIDsUser(&database, userID)
	permissionIds := PermissionIDs(&database).Where("name IN ?", permissionName)

	database.Table(TableName).Where("permission_id IN (?)", permissionIds).Where("role_id IN (?)", roleIds).First(&rolePermission)

	return rolePermission
}

// Create using for rolePermission
func Create(roleID int, permissionID int) (RolePermission, error) {
	database := *conn.DbOrm

	rolePermission := RolePermission{RoleID: roleID, PermissionID: permissionID}
	roleExisting := RolePermission{}
	isExists := database.Table(TableName).Where("role_id = ? AND permission_id = ?", roleID, permissionID).Find(&rolePermission).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("Role RolePermission is exists")
	}

	database.Table(TableName).Create(&rolePermission)
	return rolePermission, nil
}
