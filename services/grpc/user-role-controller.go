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

	if isPermited := usecases.CheckPermissionUser(data["id"].(int), []string{
		viper.GetString("permission.user-role.retrieve"),
	}); !isPermited {
		if usecases.AuthorizationValidation(token) != nil {
			return &pb.GetUserRoleResponse{}, nil
		}
	}

	userRoles := usecases.GetUserRoleWithUserID(data["id"].(int))
	if cap(userRoles) < 1 {
		return &pb.GetUserRoleResponse{}, nil
	}

	return &pb.GetUserRoleResponse{
		Items: userRoles.ToProto(),
	}, nil
}
