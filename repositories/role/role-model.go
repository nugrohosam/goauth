package role

import (
	"fmt"
	"sync"

	"github.com/fatih/structs"

	"gorm.io/gorm"
)

// TableName ...
const TableName = "roles"

// Role struct
type Role struct {
	ID   int
	Name string
}

// Roles using for many roles
type Roles []Role

// BeforeCreate ..
func (role *Role) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("beforeCreate Called")
	return nil
}

// AfterCreate ..
func (role *Role) AfterCreate(tx *gorm.DB) error {
	fmt.Println("afterCreate Called")
	return nil
}

// ToMap ..
func (roles *Roles) ToMap() []interface{} {
	var wg sync.WaitGroup

	rolesMapped := make([]interface{}, cap(*roles))
	for index, value := range *roles {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			rolesMapped[index] = structs.Map(value)
		}(&wg)
	}

	wg.Wait()

	return rolesMapped
}

// PluckName ..
func (roles *Roles) PluckName() []string {
	var wg sync.WaitGroup

	nameRolesMapped := make([]string, cap(*roles))
	for index, value := range *roles {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()
			nameRolesMapped[index] = value.Name
		}(index, &wg)
	}
	wg.Wait()

	return nameRolesMapped
}
