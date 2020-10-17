package grpc

import (
	"fmt"
	"net"

	"github.com/spf13/viper"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/nugrohosam/gosampleapi/services/grpc/pb"
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

func (newServer *server) Get(context context.Context, request *pb.Request) (*pb.Response, error) {
	token := request.GetToken()

	fmt.Println("token passed : ", token)

	return &pb.Response{
		Name:     "a",
		Username: "b",
		Email:    "c",
	}, nil
}
