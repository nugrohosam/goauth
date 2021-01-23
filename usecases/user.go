package usecases

import (
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	permissionRepo "github.com/nugrohosam/gosampleapi/repositories/permission"
	rolePermissionRepo "github.com/nugrohosam/gosampleapi/repositories/rolepermission"
	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
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
	return rolePermissionRepo.GetPermissions(ID)
}
