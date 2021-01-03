package factories

import (
	roleRepo "github.com/nugrohosam/gosampleapi/repositories/role"
)

// CreateRole ...
func CreateRole() roleRepo.Role {
	role := roleRepo.Role{
		Name: "admin",
	}

	return role
}
