package usecases

import (
	permissionRepo "github.com/nugrohosam/gosampleapi/repositories/permission"
	rolePermissionRepo "github.com/nugrohosam/gosampleapi/repositories/rolepermission"
)

// IsHaveAnyPermission ...
func IsHaveAnyPermission(userID int, permissionName []string) (bool, error) {
	isExist := rolePermissionRepo.IsExistsByUserIDAndPermissionName(userID, permissionName)
	return isExist, nil
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
