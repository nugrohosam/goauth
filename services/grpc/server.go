package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/nugrohosam/gosampleapi/services/grpc/pb"
	usecases "github.com/nugrohosam/gosampleapi/usecases"
)

// NewServer ...
var NewServer *grpc.Server

// NewServer ...
type getServiceServer struct{}

type validationServiceServer struct{}

// Listen ...
var Listen net.Listener

// Serve ...
func Serve() error {
	Prepare()

	fmt.Println("gRPC is Start to listen")
	reflection.Register(NewServer)
	if err := NewServer.Serve(Listen); err != nil {
		return err
	}

	return nil
}

// Prepare ...
func Prepare() {
	port := viper.GetString("grpc.port")
	listen, err := net.Listen("tcp", ":"+port)
	Listen = listen

	if err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}

	NewServer = grpc.NewServer()
	pb.RegisterGetServiceServer(NewServer, &getServiceServer{})
	pb.RegisterValidationServiceServer(NewServer, &validationServiceServer{})
}

// Get ...
func (getServiceServer *getServiceServer) Get(context context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	token := request.GetToken()

	if usecases.AuthorizationValidation(token) != nil {
		return &pb.GetResponse{
			Username: "",
			Name:     "",
			Email:    "",
			Id:       "",
		}, nil
	}

	data, err := usecases.GetDataAuth(token)
	if err != nil {
		return &pb.GetResponse{
			Username: "",
			Name:     "",
			Email:    "",
			Id:       "",
		}, nil
	}

	return &pb.GetResponse{
		Username: data["username"].(string),
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Id:       data["id"].(string),
	}, nil
}

// GetID ...
func (getServiceServer *getServiceServer) GetID(context context.Context, request *pb.GetRequest) (*pb.GetIdResponse, error) {
	token := request.GetToken()

	if usecases.AuthorizationValidation(token) != nil {
		return &pb.GetIdResponse{
			Id: "",
		}, nil
	}

	data, err := usecases.GetDataAuth(token)
	if err != nil {
		return &pb.GetIdResponse{
			Id: "",
		}, nil
	}

	return &pb.GetIdResponse{
		Id: data["id"].(string),
	}, nil
}

// Valdate ...
func (validationServiceServer *validationServiceServer) Validate(context context.Context, request *pb.GetRequest) (*pb.ValidationResponse, error) {
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
