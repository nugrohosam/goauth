package usecases

import (
	roleRepo "github.com/nugrohosam/gosampleapi/repositories/role"
	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
)

// IsHaveAnyRole ...
func IsHaveAnyRole(userID int, roleName []string) (bool, error) {
	isExist := userRoleRepo.IsExistsByUserIDAndRoleName(userID, roleName)
	return isExist, nil
}

// CreateRole ...
func CreateRole(name string) error {
	_, err := roleRepo.Create(name)
	return err
}

// UpdateRole ...
func UpdateRole(ID string, name string) error {
	_, err := roleRepo.Update(ID, name)
	return err
}
