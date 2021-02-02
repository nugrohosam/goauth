package permission

import (
	"fmt"
	"sync"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

// TableName ...
const TableName = "permissions"

// Permission struct
type Permission struct {
	ID   int
	Name string
}

// Permissions using for many permissions
type Permissions []Permission

// BeforeCreate ..
func (permission *Permission) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("beforeCreate Called")
	return nil
}

// AfterCreate ..
func (permission *Permission) AfterCreate(tx *gorm.DB) error {
	fmt.Println("afterCreate Called")
	return nil
}

// ToMap ..
func (permissions *Permissions) ToMap() []interface{} {
	var wg sync.WaitGroup
	wg.Add(cap(*permissions))

	permissionsMapped := make([]interface{}, cap(*permissions))
	for index, value := range *permissions {
		go func() {
			defer wg.Done()
			permissionsMapped[index] = structs.Map(value)
		}()
	}

	wg.Wait()

	return permissionsMapped
}

// PluckName ..
func (permissions *Permissions) PluckName() []string {
	var wg sync.WaitGroup
	wg.Add(cap(*permissions))

	namePermissionsMapped := make([]string, cap(*permissions))
	for index, value := range *permissions {
		go func() {
			defer wg.Done()
			namePermissionsMapped[index] = value.Name
		}()
	}

	wg.Wait()

	return namePermissionsMapped
}
