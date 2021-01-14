package rolepermission

import (
	"errors"
	"strconv"
	"strings"

	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Get using for rolePermission
func Get(search, limit, offset, orderBy string) (RolePermissions, int, error) {
	var rolePermission = RolePermissions{}
	database := *conn.DbOrm

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	totalRows := database.Table(TableName).Preload("Role", "name like ?", "%"+search+"%").Find(&rolePermission).RowsAffected
	database.Table(TableName).Preload("Role", "name like ?", "%"+search+"%").Limit(limitInt).Offset(offsetInt).Order("id " + orderBy).Find(&rolePermission)

	return rolePermission, int(totalRows), nil
}

// FindWithID is using
func FindWithID(ID string) RolePermission {
	database := *conn.DbOrm

	rolePermission := RolePermission{}
	database.Table(TableName).Where("id = ?", ID).First(&rolePermission)

	return rolePermission
}

// FindByUserIDAndPermissionName is using
func FindByUserIDAndPermissionName(userID string, permissionName []string) RolePermission {
	database := *conn.DbOrm

	roleRolePermission := RolePermission{}
	userRoles := userRoleRepo.GetByUserID(userID)
	rolesIds := userRoleRepo.PluckRolesID(userRoles)

	database.Table(TableName).Preload("Permission", "name IN (?)", strings.Join(permissionName, ",")).Where("role_id IN (?)", strings.Join(rolesIds, ",")).First(&roleRolePermission)

	return roleRolePermission
}

// IsExistsByUserIDAndPermissionName is using
func IsExistsByUserIDAndPermissionName(userID string, permissionName []string) bool {
	data := FindByUserIDAndPermissionName(userID, permissionName)
	return data.Permission.ID > 0
}

// Create using for roleRolePermission
func Create(roleID, permissionID string) (RolePermission, error) {
	database := *conn.DbOrm

	roleIDInt, _ := strconv.Atoi(roleID)
	permissionIDInt, _ := strconv.Atoi(permissionID)

	roleRolePermission := RolePermission{RoleID: roleIDInt, PermissionID: permissionIDInt}
	roleExisting := RolePermission{}
	isExists := database.Table(TableName).Where("role_id = ? AND permission_id = ?", roleRolePermission.RoleID, roleRolePermission.PermissionID).Find(&roleRolePermission).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("Role RolePermission is exists")
	}

	database.Table(TableName).Create(&roleRolePermission)
	return roleRolePermission, nil
}
