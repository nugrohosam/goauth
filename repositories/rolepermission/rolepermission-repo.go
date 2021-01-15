package rolepermission

import (
	"errors"
	"strconv"
	"strings"

	roleRepo "github.com/nugrohosam/gosampleapi/repositories/role"
	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
	conn "github.com/nugrohosam/gosampleapi/services/databases"
	"gorm.io/gorm/clause"
)

// Get using for rolePermission
func Get(search, limit, offset, orderBy string) (RolePermissions, int, error) {
	var rolePermissions = RolePermissions{}
	database := *conn.DbOrm

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	rolesSearchNameSubQuery := database.Table(roleRepo.TableName).Select("count(id)").Where("name like ?", "%"+search+"%")

	totalRows := database.Table(TableName).Where("0 < ?", rolesSearchNameSubQuery).Find(&RolePermissions{}).RowsAffected
	database.Table(TableName).Where("0 < (?)", rolesSearchNameSubQuery).Limit(limitInt).Offset(offsetInt).Order("id " + orderBy).Find(&rolePermissions)

	return rolePermissions, int(totalRows), nil
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
