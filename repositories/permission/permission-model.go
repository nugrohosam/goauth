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
	gorm.Model
	ID   int `gorm:"primaryKey;autoIncrement:true"`
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

	permissionsMapped := make([]interface{}, cap(*permissions))
	for index, value := range *permissions {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			permissionsMapped[index] = structs.Map(value)
		}(&wg)
	}

	wg.Wait()

	return permissionsMapped
}

// PluckName ..
func (permissions *Permissions) PluckName() []string {
	var wg sync.WaitGroup

	namePermissionsMapped := make([]string, cap(*permissions))
	for index, value := range *permissions {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()
			namePermissionsMapped[index] = value.Name
		}(index, &wg)
	}
	wg.Wait()

	return namePermissionsMapped
}
