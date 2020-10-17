package main

import (
	"flag"

	database "github.com/nugrohosam/gosampleapi/services/databases"
	grpcConn "github.com/nugrohosam/gosampleapi/services/grpc"
	httpConn "github.com/nugrohosam/gosampleapi/services/http"
	viper "github.com/spf13/viper"
)

func main() {
	stage := flag.String("stage", "dev", "-")

	flag.Parse()

	// initial call to envinronment variable
	if *stage == "prod" {
		viper.SetConfigFile(".env.prod.yaml")
	} else {
		viper.SetConfigFile(".env.yaml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err := database.Conn(); err != nil {
		panic(err)
	}

	if err := grpcConn.Serve(); err != nil {
		panic(err)
	}

	if err := httpConn.Serve(); err != nil {
		panic(err)
	}
}
