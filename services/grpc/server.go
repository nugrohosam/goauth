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

type server struct{}

// Serve ...
func Serve() error {
	port := viper.GetString("grpc.port")
	listen, err := net.Listen("tcp", ":"+port)

	fmt.Println("gRPC is Start to listen")
	if err != nil {
		return err
	}

	newServer := grpc.NewServer()
	pb.RegisterGetServiceServer(newServer, &server{})

	reflection.Register(newServer)
	if err := newServer.Serve(listen); err != nil {
		return err
	}

	return nil
}

// TestServe ...
func TestServe() {
	port := viper.GetString("grpc.port")
	listen, err := net.Listen("tcp", ":"+port)

	fmt.Println("gRPC is Start to listen")
	if err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}

	newServer := grpc.NewServer()
	pb.RegisterGetServiceServer(newServer, &server{})

	go func() {
		if err := newServer.Serve(listen); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func (newServer *server) Get(context context.Context, request *pb.Request) (*pb.Response, error) {
	token := request.GetToken()

	if usecases.AuthorizationValidation(token) != nil {
		return &pb.Response{
			Username: "",
			Name:     "",
			Email:    "",
		}, nil
	}

	data, err := usecases.GetDataAuth(token)
	if err != nil {
		return &pb.Response{
			Username: "",
			Name:     "",
			Email:    "",
		}, nil
	}

	return &pb.Response{
		Username: data["username"].(string),
		Name:     data["name"].(string),
		Email:    data["email"].(string),
	}, nil
}
