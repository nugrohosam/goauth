package permission

import (
	"errors"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Create using for permission
func Create(name string) (Permission, error) {
	database := *conn.Db

	permission := Permission{Name: name}
	permissionExisting := Permission{}
	isExists := database.Where("name = ?", permission.Name).Find(&permission).RowsAffected

	if isExists != 0 {
		return permissionExisting, errors.New("Permission is exists")
	}

	database.Create(&permission)
	return permission, nil
}

// Find is using
func Find(id string) Permission {
	database := *conn.Db

	permission := Permission{}
	database.Where("id = ?", id).First(&permission)

	return permission
}

// FindByEmailOrPermissionname ...
func FindByEmailOrPermissionname(emailOrPermissionname string) Permission {
	database := *conn.Db

	permission := Permission{}
	database.Where("permissionname = ? OR email = ?", emailOrPermissionname, emailOrPermissionname).First(&permission)

	return permission
}
