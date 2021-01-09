package rolepermission

import (
	"errors"
	"strconv"
	"strings"

	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Get using for permission
func Get(search, limit, offset, orderBy string) (RolePermissions, int, error) {
	var permissions = RolePermissions{}
	database := *conn.DbOrm

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	totalRows := database.Where("name like ?", "%"+search+"%").Find(&permissions).RowsAffected
	database.Where("name like ?", "%"+search+"%").Limit(limitInt).Offset(offsetInt).Order("name " + orderBy).Find(&permissions)

	return permissions, int(totalRows), nil
}

// FindWithID is using
func FindWithID(ID string) RolePermission {
	database := *conn.DbOrm

	permission := RolePermission{}
	database.Where("id = ?", ID).First(&permission)

	return permission
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
	isExists := database.Where("role_id = ? AND permission_id = ?", roleRolePermission.RoleID, roleRolePermission.PermissionID).Find(&roleRolePermission).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("Role RolePermission is exists")
	}

	database.Create(&roleRolePermission)
	return roleRolePermission, nil
}
