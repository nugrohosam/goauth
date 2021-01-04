package usecases

import (
	rolePermissionRepo "github.com/nugrohosam/gosampleapi/repositories/rolepermission"
)

// IsHaveAnyPermission ...
func IsHaveAnyPermission(userID string, permissionName []string) (bool, error) {
	isExist := rolePermissionRepo.IsExistsByUserIDAndPermissionName(userID, permissionName)
	return isExist, nil
}
