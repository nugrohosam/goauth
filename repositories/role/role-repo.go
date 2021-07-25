package role

import (
	"errors"
	"strconv"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Get using for permission
func Get(search, limit, offset, orderBy string) (Roles, int, error) {
	var roles = Roles{}
	database := *conn.DbOrm

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	totalRows := database.Where("name LIKE ?", "%"+search+"%").Find(&roles).RowsAffected
	database.Where("name LIKE ?", "%"+search+"%").Limit(limitInt).Offset(offsetInt).Order("name " + orderBy).Find(&roles)

	return roles, int(totalRows), nil
}

// Create using for role
func Create(name string) (Role, error) {
	database := *conn.DbOrm

	role := Role{Name: name}
	roleExisting := Role{}
	isExists := database.Where("name = ?", role.Name).Find(&role).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("role is exists")
	}

	database.Create(&role)
	return role, nil
}

// FindWithID is using
func FindWithID(ID int) Role {
	database := *conn.DbOrm

	role := Role{}
	database.Where("id = ?", ID).First(&role)

	return role
}

// Update using for role
func Update(ID int, name string) (Role, error) {
	database := *conn.DbOrm

	role := Role{Name: name}
	roleExisting := Role{}
	isExists := database.Where("name = ?", role.Name).Where("id != ?", ID).Find(&role).RowsAffected

	if isExists != 0 {
		return roleExisting, errors.New("role is exists")
	}

	database.Model(Role{}).Where("id = ?", ID).Updates(&role)
	return role, nil
}

// Delete is using
func Delete(ID int) {
	database := *conn.DbOrm
	database.Delete(&Role{}, ID)
}

// Find is using
func Find(ID int) Role {
	database := *conn.DbOrm

	role := Role{}
	database.Where("id = ?", ID).First(&role)

	return role
}
