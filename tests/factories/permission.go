package factories

import (
	permissionRepo "github.com/nugrohosam/gosampleapi/repositories/permission"
)

// CreatePermission ...
func CreatePermission() permissionRepo.Permission {
	permission := permissionRepo.Permission{
		Name: "open-cashier",
	}

	return permission
}
