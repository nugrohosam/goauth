package usecases

import (
	permissionRepo "github.com/nugrohosam/gosampleapi/repositories/permission"
	rolePermissionRepo "github.com/nugrohosam/gosampleapi/repositories/rolepermission"
	userRepo "github.com/nugrohosam/gosampleapi/repositories/user"
	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
)

// ShowUser ..
func ShowUser(ID string) userRepo.User {
	user := userRepo.Find(ID)
	return user
}

// GetUserPermissionWithUserID ..
func GetUserPermissionWithUserID(ID string) permissionRepo.Permissions {
	userRoles := userRoleRepo.GetByUserID(ID)
	roleIDs := userRoles.PluckRoleID()

	rolePermissionRepo.GetByRoleIDs(roleIDs)
	return permissionRepo.Permissions{}
}
