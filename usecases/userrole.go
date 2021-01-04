package usecases

import (
	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
)

// IsHaveAnyRole ...
func IsHaveAnyRole(userID string, roleName []string) (bool, error) {
	isExist := userRoleRepo.IsExistsByUserIDAndRoleName(userID, roleName)
	return isExist, nil
}

// CreateUserRole ...
func CreateUserRole(userID, roleID string) error {
	_, err := userRoleRepo.Create(userID, roleID)
	return err
}

// UpdateUserRole ...
func UpdateUserRole(ID, roleID, userID string) error {
	_, err := userRoleRepo.Update(ID, roleID, userID)
	return err
}

// DeleteUserRole ...
func DeleteUserRole(ID string) error {
	userRoleRepo.Delete(ID)
	return nil
}
