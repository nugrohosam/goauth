package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nugrohosam/gosampleapi/services/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGetServiceClient(conn)

	// Contact the server and prnt out its riesponse.
	ctx := context.Background()
	token := "Anjeng"
	req := &pb.Request{Token: token}
	res, err := client.Get(ctx, req)

	fmt.Println(
		"username :", res.Username,
		"email :", res.Email,
		"name :", res.Name,
	)
}
