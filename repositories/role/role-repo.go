package role

import (
	"errors"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Create using for role
func Create(name string) (Role, error) {
	database := *conn.DbOrm

	role := Role{Name: name}
	roleExisting := Role{}
	isExists := database.Where("name = ?", role.Name).Find(&role).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("Role is exists")
	}

	database.Create(&role)
	return role, nil
}

// Update using for role
func Update(ID string, name string) (Role, error) {
	database := *conn.DbOrm

	role := Role{Name: name}
	roleExisting := Role{}
	isExists := database.Where("name = ?", role.Name).Where("id != ?", ID).Find(&role).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("Role is exists")
	}

	database.Model(Role{}).Where("id = ?", ID).Updates(&role)
	return role, nil
}

// Delete is using
func Delete(ID string) {
	database := *conn.DbOrm
	database.Delete(&Role{}, ID)
}

// Find is using
func Find(id string) Role {
	database := *conn.DbOrm

	role := Role{}
	database.Where("id = ?", id).First(&role)

	return role
}
