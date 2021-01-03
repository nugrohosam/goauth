package userrole

import (
	"errors"
	"strings"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Create using for userRole
func Create(roleID, userID int) (UserRole, error) {
	database := *conn.Db

	userRole := UserRole{RoleID: roleID, UserID: userID}
	roleExisting := UserRole{}
	isExists := database.Table(TableName).Where("role_id = ? AND user_id = ?", userRole.RoleID, userRole.UserID).Find(&userRole).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("Role Permission is exists")
	}

	database.Create(&userRole)
	return userRole, nil
}

// Find is using
func Find(id string) UserRole {
	database := *conn.Db

	userRole := UserRole{}
	database.Where("id = ?", id).First(&userRole)

	return userRole
}

// FindByUserIDAndRoleName is using
func FindByUserIDAndRoleName(userID int, roleName []string) UserRole {
	database := *conn.Db

	rolePermission := UserRole{}
	database.Table(TableName).Preload("Role", "name IN (?)", strings.Join(roleName, ",")).Where("user_id = ?", userID).First(&rolePermission)

	return rolePermission
}

// IsExistsByUserIDAndRoleName is using
func IsExistsByUserIDAndRoleName(userID int, roleName []string) bool {
	data := FindByUserIDAndRoleName(userID, roleName)
	return data.ID > 0
}
