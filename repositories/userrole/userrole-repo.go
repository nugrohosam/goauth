package userrole

import (
	"errors"
	"strconv"
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

	userRole := UserRole{}
	database.Table(TableName).Preload("Role", "name IN (?)", strings.Join(roleName, ",")).Where("user_id = ?", userID).First(&userRole)

	return userRole
}

// GetByUserID is using
func GetByUserID(userID int) []UserRole {
	database := *conn.Db

	userRoles := []UserRole{}
	database.Table(TableName).Where("user_id = ?", userID).Find(&userRoles)

	return userRoles
}

// GetRolesID is using
func GetRolesID(userRoles []UserRole) []string {

	i := 0
	lengthUserRoles := cap(userRoles)
	mapped := make([]string, lengthUserRoles)

	for _, userRole := range userRoles {
		mapped[i] = strconv.Itoa(userRole.RoleID)
		i++
	}

	return mapped
}

// IsExistsByUserIDAndRoleName is using
func IsExistsByUserIDAndRoleName(userID int, roleName []string) bool {
	data := FindByUserIDAndRoleName(userID, roleName)
	return data.ID > 0
}
