package usecases

import (
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	permissionRepo "github.com/nugrohosam/gosampleapi/repositories/permission"
	"github.com/spf13/viper"
)

// CheckPermissionUser ..
func CheckPermissionUser(userID string, permissionAccess []string) bool {

	globalAccessOpen := viper.GetString("permission.global.open")
	globalAccessClose := viper.GetString("permission.global.close")

	if isExists, err := IsHaveAnyPermission(userID, []string{globalAccessOpen}); isExists || err != nil {
		return true
	} else if isExists, err := IsHaveAnyPermission(userID, []string{globalAccessClose}); isExists || err != nil {
		return false
	} else if isExists, err := IsHaveAnyPermission(userID, permissionAccess); !isExists || err != nil {
		return false
	}

	return true
}

// GetPermission ...
func GetPermission(serach, perPage, page, order string) ([]permissionRepo.Permission, int, error) {

	availableOrder := map[string]string{
		"atoz": "asc",
		"ztoa": "desc",
	}

	orderBy := availableOrder[order]
	limit, offset := helpers.GenerateLimitOffset(perPage, page)

	permissions, total, err := permissionRepo.Get(serach, limit, offset, orderBy)
	if err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}

// ShowPermission ...
func ShowPermission(ID string) permissionRepo.Permission {
	permission := permissionRepo.FindWithID(ID)

	return permission
}

// CreatePermission ...
func CreatePermission(name string) error {
	_, err := permissionRepo.Create(name)
	return err
}

// UpdatePermission ...
func UpdatePermission(ID string, name string) error {
	_, err := permissionRepo.Update(ID, name)
	return err
}

// DeletePermission ...
func DeletePermission(ID string) error {
	permissionRepo.Delete(ID)
	return nil
}
