package usecases

import (
	"fmt"

	helpers "github.com/nugrohosam/gosampleapi/helpers"
	rolePermissionRepo "github.com/nugrohosam/gosampleapi/repositories/rolepermission"
)

// IsHaveAnyPermission ...
func IsHaveAnyPermission(userID string, permissionName []string) (bool, error) {
	rolePermission := rolePermissionRepo.FindByUserIDAndPermissionName(userID, permissionName)
	return (rolePermission.ID > 0), nil
}

// GetRolePermission ...
func GetRolePermission(search, perPage, page, order string) (rolePermissionRepo.RolePermissions, int, error) {
	availableOrder := map[string]string{
		"atoz": "asc",
		"ztoa": "desc",
	}

	orderBy := availableOrder[order]
	limit, offset := helpers.GenerateLimitOffset(perPage, page)

	rolePermissions, total, err := rolePermissionRepo.Get(search, limit, offset, orderBy)
	fmt.Println(rolePermissions)
	if err != nil {
		return nil, 0, err
	}

	return rolePermissions, total, nil
}

// ShowRolePermission ...
func ShowRolePermission(ID string) rolePermissionRepo.RolePermission {
	permission := rolePermissionRepo.FindWithID(ID)

	return permission
}

// CreateRolePermission ...
func CreateRolePermission(roleID int, permissionID int) error {
	_, err := rolePermissionRepo.Create(roleID, permissionID)
	return err
}
