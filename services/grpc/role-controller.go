package grpc

import (
	"context"
	"strconv"

	pb "github.com/nugrohosam/gosampleapi/services/grpc/pb"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
	"github.com/spf13/viper"
)

type getRoleServiceServer struct{}

// Get ...
func (getRoleServiceServer *getRoleServiceServer) GetRoleWithID(context context.Context, request *pb.GetRoleRequest) (*pb.GetRoleResponse, error) {
	token := request.GetToken()

	if usecases.AuthorizationValidation(token) != nil {
		return &pb.GetRoleResponse{}, nil
	}

	data, _ := usecases.GetDataAuth(token)

	if isPermited := usecases.CheckPermissionUser(data["id"].(int), []string{
		viper.GetString("permission.role.retrieve"),
	}); !isPermited {
		if usecases.AuthorizationValidation(token) != nil {
			return &pb.GetRoleResponse{}, nil
		}
	}

	roleID := request.GetRoleId()
	roleIDInt, _ := strconv.Atoi(roleID)

	role := usecases.ShowRole(roleIDInt)
	if role.ID < 1 {
		return &pb.GetRoleResponse{}, nil
	}

	return role.ToProto(), nil
}
