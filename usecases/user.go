package usecases

import (
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	permissionRepo "github.com/nugrohosam/gosampleapi/repositories/permission"
	roleRepo "github.com/nugrohosam/gosampleapi/repositories/role"
	rolePermissionRepo "github.com/nugrohosam/gosampleapi/repositories/rolepermission"
	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
)

// GetUser ..
func GetUser(search, perPage, page, order string) (userRepo.Users, int, error) {
	availableOrder := map[string]string{
		"atoz": "asc",
		"ztoa": "desc",
	}

	orderBy := availableOrder[order]
	limit, offset := helpers.GenerateLimitOffset(perPage, page)

	users, total, err := userRepo.Get(search, limit, offset, orderBy)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// ShowUser ..
func ShowUser(ID int) userRepo.User {
	user := userRepo.Find(ID)
	return user
}

// GetPermissionWithUserID ..
func GetPermissionWithUserID(ID int) permissionRepo.Permissions {
	permissions := rolePermissionRepo.GetPermissionsWithUserID(ID)
	return permissions
}

// GetRoleWithUserID ..
func GetRoleWithUserID(ID int) roleRepo.Roles {
	roles := userRoleRepo.GetRolesWithUserID(ID)
	return roles
}

// IsHaveAnyPermission ...
func IsHaveAnyPermission(userID int, permissionName []string) (bool, error) {
	permissions := GetPermissionWithUserID(userID)
	userPermissionsName := permissions.PluckName()

	isTrue := false
	for _, value := range permissionName {
		isTrue = helpers.StringInSlice(value, userPermissionsName)
		if isTrue {
			break
		}
	}

	return isTrue, nil
}
