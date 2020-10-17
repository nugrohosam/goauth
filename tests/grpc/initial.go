package grpc

import (
	"testing"

	database "github.com/nugrohosam/gosampleapi/services/databases"
	grpcServe "github.com/nugrohosam/gosampleapi/services/grpc"
	viper "github.com/spf13/viper"
)

// InitialTest ...
func InitialTest(t *testing.T) {
	// initial call to envinronment variable
	viper.SetConfigType("yaml")
	viper.SetConfigFile("../../.env.test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Error(err.Error())
		panic(err)
	}

	if err := database.Conn(); err != nil {
		t.Error(err.Error())
		panic(err)
	}

	grpcServe.TestServe()
}
