package usecases

import (
	userRoleRepo "github.com/nugrohosam/gosampleapi/repositories/userrole"
)

// IsHaveAnyRole ...
func IsHaveAnyRole(userID int, roleName []string) (bool, error) {
	isExist := userRoleRepo.IsExistsByUserIDAndRoleName(userID, roleName)
	return isExist, nil
}
