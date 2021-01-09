package usecases

import (
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	rolePermissionRepo "github.com/nugrohosam/gosampleapi/repositories/rolepermission"
)

// IsHaveAnyPermission ...
func IsHaveAnyPermission(userID string, permissionName []string) (bool, error) {
	isExist := rolePermissionRepo.IsExistsByUserIDAndPermissionName(userID, permissionName)
	return isExist, nil
}

// GetRolePermission ...
func GetRolePermission(serach, perPage, page, order string) ([]rolePermissionRepo.RolePermission, int, error) {
	availableOrder := map[string]string{
		"atoz": "asc",
		"ztoa": "desc",
	}

	orderBy := availableOrder[order]
	limit, offset := helpers.GenerateLimitOffset(perPage, page)

	permissions, total, err := rolePermissionRepo.Get(serach, limit, offset, orderBy)
	if err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}

// ShowRolePermission ...
func ShowRolePermission(ID string) rolePermissionRepo.RolePermission {
	permission := rolePermissionRepo.FindWithID(ID)

	return permission
}

// CreateRolePermission ...
func CreateRolePermission(roleID, permissionID string) error {
	_, err := rolePermissionRepo.Create(roleID, permissionID)
	return err
}
