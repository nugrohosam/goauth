package main

import (
	"bufio"
	"flag"
	"os"

	database "github.com/nugrohosam/gosampleapi/services/databases"
	grpcConn "github.com/nugrohosam/gosampleapi/services/grpc"
	httpConn "github.com/nugrohosam/gosampleapi/services/http"
	infrastructure "github.com/nugrohosam/gosampleapi/services/insfrastructure"
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

	infrastructure.PrepareSentry()

	if err := database.Conn(); err != nil {
		panic(err)
	}

	runGrpc := func() {
		if err := grpcConn.Serve(); err != nil {
			panic(err)
		}
	}

	runHTTP := func() {
		if err := httpConn.Serve(); err != nil {
			panic(err)
		}
	}

	go runGrpc()
	go runHTTP()

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
