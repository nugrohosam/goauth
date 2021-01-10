package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/nugrohosam/gosampleapi/services/grpc/pb"
)

// NewServer ...
var NewServer *grpc.Server

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
	pb.RegisterGetAuthServiceServer(NewServer, &getAuthServiceServer{})
	pb.RegisterValidationServiceServer(NewServer, &validationServiceServer{})
	pb.RegisterGetRoleServiceServer(NewServer, &getRoleServiceServer{})
	pb.RegisterGetUserRoleServiceServer(NewServer, &getUserRoleServiceServer{})
}
