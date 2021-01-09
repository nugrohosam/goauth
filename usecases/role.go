package usecases

import (
	helpers "github.com/nugrohosam/gosampleapi/helpers"
	roleRepo "github.com/nugrohosam/gosampleapi/repositories/role"
)

// GetRole ...
func GetRole(serach, perPage, page, order string) ([]roleRepo.Role, int, error) {

	availableOrder := map[string]string{
		"atoz": "asc",
		"ztoa": "desc",
	}

	orderBy := availableOrder[order]
	limit, offset := helpers.GenerateLimitOffset(perPage, page)

	roles, total, err := roleRepo.Get(serach, limit, offset, orderBy)
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

// ShowRole ...
func ShowRole(ID string) roleRepo.Role {
	role := roleRepo.FindWithID(ID)

	return role
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

// DeleteRole ...
func DeleteRole(ID string) error {
	roleRepo.Delete(ID)
	return nil
}
