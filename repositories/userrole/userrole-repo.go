package userrole

import (
	"errors"
	"strconv"
	"strings"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Create using for userRole
func Create(roleID, userID string) (UserRole, error) {
	database := *conn.DbOrm

	roleIDInt, _ := strconv.Atoi(roleID)
	userIDInt, _ := strconv.Atoi(userID)

	userRole := UserRole{RoleID: roleIDInt, UserID: userIDInt}
	roleExisting := UserRole{}
	isExists := database.Table(TableName).Where("role_id = ? AND user_id = ?", userRole.RoleID, userRole.UserID).Find(&userRole).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("Role Permission is exists")
	}

	database.Create(&userRole)
	return userRole, nil
}

// Update using for userRole
func Update(ID, roleID, userID string) (UserRole, error) {
	database := *conn.DbOrm

	roleIDInt, _ := strconv.Atoi(roleID)
	userIDInt, _ := strconv.Atoi(userID)

	userRole := UserRole{RoleID: roleIDInt, UserID: userIDInt}
	roleExisting := UserRole{}
	isExists := database.Table(TableName).Where("role_id = ? AND user_id = ?", userRole.RoleID, userRole.UserID).Where("id != ?", ID).Find(&userRole).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("User is has this role")
	}

	database.Table(TableName).Model(UserRole{}).Where("id = ?", ID).Updates(&userRole)
	return userRole, nil
}

// Find is using
func Find(id string) UserRole {
	database := *conn.DbOrm

	userRole := UserRole{}
	database.Where("id = ?", id).First(&userRole)

	return userRole
}

// FindByUserIDAndRoleName is using
func FindByUserIDAndRoleName(userID string, roleName []string) UserRole {
	database := *conn.DbOrm

	userRole := UserRole{}
	database.Table(TableName).Preload("Role", "name IN (?)", strings.Join(roleName, ",")).Where("user_id = ?", userID).First(&userRole)

	return userRole
}

// GetByUserID is using
func GetByUserID(userID string) []UserRole {
	database := *conn.DbOrm

	userRoles := []UserRole{}
	database.Table(TableName).Where("user_id = ?", userID).Find(&userRoles)

	return userRoles
}

// PluckRolesID is using
func PluckRolesID(userRoles []UserRole) []string {

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
func IsExistsByUserIDAndRoleName(userID string, roleName []string) bool {
	data := FindByUserIDAndRoleName(userID, roleName)
	return data.ID > 0
}

// Delete is using
func Delete(ID string) {
	database := *conn.DbOrm
	database.Delete(&UserRole{}, ID)
}
