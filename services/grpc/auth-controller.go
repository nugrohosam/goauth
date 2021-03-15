package grpc

import (
	"context"

	pb "github.com/nugrohosam/gosampleapi/services/grpc/pb"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

type getAuthServiceServer struct{}

type validationServiceServer struct{}

// Get ...
func (getAuthServiceServer *getAuthServiceServer) GetAuth(context context.Context, request *pb.GetAuthRequest) (*pb.GetAuthResponse, error) {
	token := request.GetToken()

	if usecases.AuthorizationValidation(token) != nil {
		return &pb.GetAuthResponse{}, nil
	}

	data, err := usecases.GetDataAuth(token)
	if err != nil {
		return &pb.GetAuthResponse{}, nil
	}

	return &pb.GetAuthResponse{
		Username: data["username"].(string),
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Id:       data["id"].(float64),
	}, nil
}

// GetID ...
func (getAuthServiceServer *getAuthServiceServer) GetAuthID(context context.Context, request *pb.GetAuthRequest) (*pb.GetAuthIdResponse, error) {
	token := request.GetToken()

	if usecases.AuthorizationValidation(token) != nil {
		return &pb.GetAuthIdResponse{}, nil
	}

	data, err := usecases.GetDataAuth(token)
	if err != nil {
		return &pb.GetAuthIdResponse{}, nil
	}

	return &pb.GetAuthIdResponse{
		Id: data["id"].(float64),
	}, nil
}

// Valdate ...
func (validationServiceServer *validationServiceServer) Validate(context context.Context, request *pb.GetAuthRequest) (*pb.ValidationResponse, error) {
	token := request.GetToken()

	if usecases.AuthorizationValidation(token) != nil {
		return &pb.ValidationResponse{
			Valid: false,
		}, nil
	}

	return &pb.ValidationResponse{
		Valid: true,
	}, nil
}
