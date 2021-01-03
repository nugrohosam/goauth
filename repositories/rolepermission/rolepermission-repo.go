package rolepermission

import (
	"errors"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

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
