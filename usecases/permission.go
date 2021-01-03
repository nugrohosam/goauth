package usecases

import (
	permissionRepo "github.com/nugrohosam/gosampleapi/repositories/permission"
)

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
