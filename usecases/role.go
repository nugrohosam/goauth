package usecases

import (
	roleRepo "github.com/nugrohosam/gosampleapi/repositories/role"
)

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
