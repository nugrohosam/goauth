package userrole

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	roleRepo "github.com/nugrohosam/gosampleapi/repositories/role"
	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
	conn "github.com/nugrohosam/gosampleapi/services/databases"
	"gorm.io/gorm/clause"
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
func Create(roleID, userID int) (UserRole, error) {
	database := *conn.DbOrm

	userRole := UserRole{RoleID: roleID, UserID: userID}
	roleExisting := UserRole{}
	isExists := database.Table(TableName).Where("role_id = ? AND user_id = ?", roleID, userID).Find(&userRole).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("Role Permission is exists")
	}

	database.Table(TableName).Create(&userRole)
	return userRole, nil
}

// Update using for userRole
func Update(ID, roleID, userID int) (UserRole, error) {
	database := *conn.DbOrm

	userRole := UserRole{RoleID: roleID, UserID: userID}
	roleExisting := UserRole{}
	isExists := database.Table(TableName).Where("role_id = ? AND user_id = ?", roleID, userID).Where("id != ?", ID).Find(&userRole).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("User is has this role")
	}

	database.Table(TableName).Model(UserRole{}).Where("id = ?", ID).Updates(&userRole)
	return userRole, nil
}

// Find is using
func Find(ID int) UserRole {
	database := *conn.DbOrm

	userRole := UserRole{}
	database.Table(TableName).Preload(clause.Associations).Where("id = ?", ID).First(&userRole)

	return userRole
}

// FindByUserIDAndRoleName is using
func FindByUserIDAndRoleName(userID int, roleName []string) UserRole {
	database := *conn.DbOrm

	userRole := UserRole{}
	database.Table(TableName).Preload("Role", "name IN (?)", strings.Join(roleName, ",")).Where("user_id = ?", userID).First(&userRole)

	return userRole
}

// GetByUserID is using
func GetByUserID(userID int) UserRoles {
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
		mapped[i] = fmt.Sprint(userRole)
		i++
	}

	return mapped
}

// IsExistsByUserIDAndRoleName is using
func IsExistsByUserIDAndRoleName(userID int, roleName []string) bool {
	data := FindByUserIDAndRoleName(userID, roleName)
	return data.ID > 0
}

// Delete is using
func Delete(ID int) (UserRole, error) {
	database := *conn.DbOrm

	userRole := UserRole{}
	database.Table(TableName).Delete(&userRole, ID)

	return userRole, nil
}

// GetRolesWithUserID ..
func GetRolesWithUserID(userID int) roleRepo.Roles {
	database := *conn.DbOrm

	roles := roleRepo.Roles{}

	scopeQuery := database.Table(TableName).Select("role_id").Where("user_id = ?", userID)
	database.Table(roleRepo.TableName).Where("role_id IN (?)", scopeQuery).Find(&roles)

	return roles
}

// GetUsersWithRoleID ..
func GetUsersWithRoleID(roleID int) userRepo.Users {
	database := *conn.DbOrm

	users := userRepo.Users{}

	scopeQuery := database.Table(TableName).Select("user_id").Where("role_id = ?", roleID)
	database.Table(userRepo.TableName).Where("user_id IN (?)", scopeQuery).Find(&users)

	return users
}
