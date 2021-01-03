package role

import (
	"fmt"

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
