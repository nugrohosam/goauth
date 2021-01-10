package role

import (
	"fmt"

	pb "github.com/nugrohosam/gosampleapi/services/grpc/pb"

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

func (role *Role) ToProto() (*pb.GetRoleResponse) {
	return &pb.GetRoleResponse{
		Id:   int64(role.ID),
		Name: role.Name,
	}
}