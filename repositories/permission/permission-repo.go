package permission

import (
	"errors"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Create using for permission
func Create(name string) (Permission, error) {
	database := *conn.DbOrm

	permission := Permission{Name: name}
	permissionExisting := Permission{}
	isExists := database.Where("name = ?", permission.Name).Find(&permission).RowsAffected

	if isExists != 0 {
		return permissionExisting, errors.New("Permission is exists")
	}

	database.Create(&permission)
	return permission, nil
}

// Update using for permission
func Update(ID string, name string) (Permission, error) {
	database := *conn.DbOrm

	permission := Permission{Name: name}
	permissionExisting := Permission{}
	isExists := database.Where("name = ?", permission.Name).Where("id != ?", ID).Find(&permission).RowsAffected

	if isExists != 0 {
		return permissionExisting, errors.New("Permission is exists")
	}

	database.Model(Permission{}).Where("id = ?", ID).Updates(&permission)
	return permission, nil
}

// Delete is using
func Delete(ID string) {
	database := *conn.DbOrm
	database.Delete(&Permission{}, ID)
}

// Find is using
func Find(ID string) Permission {
	database := *conn.DbOrm

	permission := Permission{}
	database.Where("id = ?", ID).First(&permission)

	return permission
}

// FindByEmailOrPermissionname ...
func FindByEmailOrPermissionname(emailOrPermissionname string) Permission {
	database := *conn.DbOrm

	permission := Permission{}
	database.Where("permissionname = ? OR email = ?", emailOrPermissionname, emailOrPermissionname).First(&permission)

	return permission
}
