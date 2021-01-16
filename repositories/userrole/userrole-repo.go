package userrole

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/nugrohosam/gosampleapi/repositories/role"
	"github.com/nugrohosam/gosampleapi/repositories/user"
	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Get using for permission
func Get(search, limit, offset, orderBy string) (UserRoles, int, error) {
	var userRoles = UserRoles{}
	database := *conn.DbOrm

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	totalRows := database.Table(TableName).Find(&userRoles).RowsAffected
	database.Table(TableName).Limit(limitInt).Offset(offsetInt).Order("id " + orderBy).Find(&userRoles)

	return userRoles, int(totalRows), nil
}

// Create using for userRole
func Create(roleID, userID interface{}) (UserRole, error) {
	database := *conn.DbOrm

	userRole := UserRole{RoleID: roleID.(role.RoleID), UserID: userID.(user.UserID)}
	roleExisting := UserRole{}
	isExists := database.Table(TableName).Where("role_id = ? AND user_id = ?", userRole.RoleID, userRole.UserID).Find(&userRole).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("Role Permission is exists")
	}

	database.Table(TableName).Create(&userRole)
	return userRole, nil
}

// Update using for userRole
func Update(ID, roleID, userID interface{}) (UserRole, error) {
	database := *conn.DbOrm

	userRole := UserRole{RoleID: roleID.(role.RoleID), UserID: userID.(user.UserID)}
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
	database.Table(TableName).Where("id = ?", id).First(&userRole)

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
func GetByUserID(userID string) UserRoles {
	database := *conn.DbOrm

	userRoles := UserRoles{}
	database.Table(TableName).Where("user_id = ?", userID).Find(&userRoles)

	return userRoles
}

// PluckRolesID is using
func PluckRolesID(userRoles []UserRole) []string {

	i := 0
	lengthUserRoles := cap(userRoles)
	mapped := make([]string, lengthUserRoles)

	for _, userRole := range userRoles {
		mapped[i] = fmt.Sprint(userRole.RoleID)
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
	database.Table(TableName).Delete(&UserRole{}, ID)
}
