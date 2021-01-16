package userrole

import (
	"github.com/nugrohosam/gosampleapi/repositories/role"
	"github.com/nugrohosam/gosampleapi/repositories/user"
	pb "github.com/nugrohosam/gosampleapi/services/grpc/pb"
)

// TableName ...
const TableName = "user_role"

// UserRole struct
type UserRole struct {
	ID     int
	UserID int
	User   user.User `gorm:"constraint:OnDelete:CASCADE;references:id"`
	RoleID int
	Role   role.Role `gorm:"constraint:OnDelete:CASCADE;references:id"`
}

// UserRoles using for many user_role
type UserRoles []UserRole

// PluckRoleID ..
func (userRoles *UserRoles) PluckRoleID() []int {
	roleIDs := make([]int, cap(*userRoles))
	for i, userRole := range *userRoles {
		roleIDs[i] = userRole.ID
	}

	return roleIDs
}

// ToProto ..
func (userRoles *UserRoles) ToProto() []*pb.UserRole {
	userRolesResource := make([]*pb.UserRole, cap(*userRoles))
	for i, userRoleItem := range *userRoles {
		userRolesResource[i] = userRoleItem.ToProto()
	}

	return userRolesResource
}

// ToProto ..
func (userRole *UserRole) ToProto() *pb.UserRole {
	return &pb.UserRole{
		Id:     int64(userRole.ID),
		UserId: int64(userRole.UserID),
		RoleId: int64(userRole.RoleID),
	}
}
