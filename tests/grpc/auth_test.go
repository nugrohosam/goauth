package grpc

import (
	"context"
	"log"
	"testing"

	pb "github.com/nugrohosam/gosampleapi/services/grpc/pb"
	utilities "github.com/nugrohosam/gosampleapi/tests/utilities"
	assert "github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

const bufSize = 1024 * 1024

// BookResponse ...
type BookResponse struct{}

// GetBook ...
func TestRun(t *testing.T) {
	InitialTest(t)
	defer utilities.DbCleaner(t)

	authWithToken(t)
}

func authWithToken(t *testing.T) {
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
	if err != nil {
		t.Error(err.Error())
	}

	assert.NotEmpty(t, res.Username)
	assert.NotEmpty(t, res.Email)
	assert.NotEmpty(t, res.Name)
}