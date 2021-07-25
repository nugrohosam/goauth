package permission

import (
	"errors"
	"strconv"

	conn "github.com/nugrohosam/gosampleapi/services/databases"
)

// Get using for permission
func Get(search, limit, offset, orderBy string) (Permissions, int, error) {
	var permissions = Permissions{}
	database := *conn.DbOrm

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	totalRows := database.Where("name LIKE ?", "%"+search+"%").Find(&permissions).RowsAffected
	database.Where("name LIKE ?", "%"+search+"%").Limit(limitInt).Offset(offsetInt).Order("name " + orderBy).Find(&permissions)

	return permissions, int(totalRows), nil
}

// Create using for permission
func Create(name string) (Permission, error) {
	database := *conn.DbOrm

	permission := Permission{Name: name}
	permissionExisting := Permission{}
	isExists := database.Where("name = ?", permission.Name).Find(&permission).RowsAffected

	if isExists != 0 {
		return permissionExisting, errors.New("permission is exists")
	}

	database.Create(&permission)
	return permission, nil
}

// Update using for permission
func Update(ID int, name string) (Permission, error) {
	database := *conn.DbOrm

	permission := Permission{Name: name}
	permissionExisting := Permission{}
	isExists := database.Where("name = ?", permission.Name).Where("id != ?", ID).Find(&permission).RowsAffected

	if isExists != 0 {
		return permissionExisting, errors.New("permission is exists")
	}

	database.Model(Permission{}).Where("id = ?", ID).Updates(&permission)
	return permission, nil
}

// Delete is using
func Delete(ID int) {
	database := *conn.DbOrm
	database.Delete(&Permission{}, ID)
}

// FindWithID is using
func FindWithID(ID int) Permission {
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
