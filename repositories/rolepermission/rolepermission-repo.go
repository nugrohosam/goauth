package rolepermission

import (
	"errors"
	"strings"

	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// FindByUserIDAndPermissionName is using
func FindByUserIDAndPermissionName(userID int, permissionName []string) RolePermission {
	database := *conn.Db

	rolePermission := RolePermission{}
	userRoles := userRoleRepo.GetByUserID(userID)
	rolesIds := userRoleRepo.GetRolesID(userRoles)

	database.Table(TableName).Preload("Permission", "name IN (?)", strings.Join(permissionName, ",")).Where("role_id IN (?)", strings.Join(rolesIds, ",")).First(&rolePermission)

	return rolePermission
}

// IsExistsByUserIDAndPermissionName is using
func IsExistsByUserIDAndPermissionName(userID int, permissionName []string) bool {
	data := FindByUserIDAndPermissionName(userID, permissionName)
	return data.ID > 0
}

// Create using for rolePermission
func Create(roleID, permissionID int) (RolePermission, error) {
	database := *conn.Db

	rolePermission := RolePermission{RoleID: roleID, PermissionID: permissionID}
	roleExisting := RolePermission{}
	isExists := database.Where("role_id = ? AND permission_id = ?", rolePermission.RoleID, rolePermission.PermissionID).Find(&rolePermission).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("Role Permission is exists")
	}

	database.Create(&rolePermission)
	return rolePermission, nil
}
