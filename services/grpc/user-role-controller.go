package grpc

import (
	"context"

	pb "github.com/nugrohosam/gosampleapi/services/grpc/pb"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
	"github.com/spf13/viper"
)

type getUserRoleServiceServer struct{}

// Get ...
func (getUserRoleServiceServer *getUserRoleServiceServer) GetUserRole(context context.Context, request *pb.GetUserRoleRequest) (*pb.GetUserRoleResponse, error) {
	token := request.GetToken()

	if usecases.AuthorizationValidation(token) != nil {
		return &pb.GetUserRoleResponse{}, nil
	}

	data, _ := usecases.GetDataAuth(token)

	if isPermited := usecases.CheckPermissionUser(data["id"].(string), []string{
		viper.GetString("permission.user-role.retrieve"),
	}); !isPermited {
		if usecases.AuthorizationValidation(token) != nil {
			return &pb.GetUserRoleResponse{}, nil
		}
	}

	userRoles := usecases.GetUserRoleWithUserID(data["id"].(string))
	if cap(userRoles) < 1 {
		return &pb.GetUserRoleResponse{}, nil
	}

	userRolesResource := make([]*pb.UserRole, cap(userRoles))
	for i, userRoleItem := range userRoles {
		userRolesResource[i] = &pb.UserRole{
			Id:     int64(userRoleItem.ID),
			UserId: int64(userRoleItem.UserID),
			RoleId: int64(userRoleItem.RoleID),
		}
	}

	return &pb.GetUserRoleResponse{
		Items: userRolesResource,
	}, nil
}
