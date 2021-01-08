package usecases

import (
	helper "github.com/nugrohosam/gosampleapi/helpers"
	permissionRepo "github.com/nugrohosam/gosampleapi/repositories/permission"
)

// GetPermission ...
func GetPermission(serach, perPage, page, order string) ([]permissionRepo.Permission, error) {

	availableOrder := map[string]string{
		"atoz": "asc",
		"ztoa": "desc",
	}

	orderBy := availableOrder[order]
	limit, offset := helper.GenerateLimitOffset(perPage, page)

	permissions, err := permissionRepo.Get(serach, limit, offset, orderBy)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

// CreatePermission ...
func CreatePermission(name string) error {
	_, err := permissionRepo.Create(name)
	return err
}

// UpdatePermission ...
func UpdatePermission(ID string, name string) error {
	_, err := permissionRepo.Update(ID, name)
	return err
}

// DeletePermission ...
func DeletePermission(ID string) error {
	permissionRepo.Delete(ID)
	return nil
}
