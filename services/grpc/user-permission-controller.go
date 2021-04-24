package grpc

import (
	"context"

	pb "github.com/nugrohosam/gosampleapi/services/grpc/proto"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

type userPermissionServiceServer struct {
	pb.UnimplementedUserPermissionServiceServer
}

// Get ...
func (userPermissionServiceServer *userPermissionServiceServer) Validate(ctx context.Context, request *pb.GetUserPermissionRequest) (*pb.GetUserPermissionResponse, error) {
	token := request.GetToken()
	permission := request.GetPermission()

	if usecases.AuthorizationValidation(token) == nil {
		return &pb.GetUserPermissionResponse{
			IsValid: false,
		}, nil
	}

	data, _ := usecases.GetDataAuth(token)
	isPermited := usecases.CheckPermissionUser(data["id"].(int), []string{permission})
	return &pb.GetUserPermissionResponse{
		IsValid: isPermited,
	}, nil
}
